/*
   Copyright [2023] [veerdone]

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package simpleredis

import (
	"syscall"

	"github.com/veerdone/simple_redis/db"
	"go.uber.org/zap"
	"golang.org/x/sys/unix"
)

type EpollListener struct {
	efd int
	sfd int
	*zap.Logger
	db *db.DB
}

func NewListener() Listener {
	l, _ := NewLogger()
	return &EpollListener{
		Logger: l,
		db:     db.NewDB(),
	}
}

func (e *EpollListener) epollInit() error {
	efd, err := unix.EpollCreate(1)
	if err != nil {
		return err
	}
	e.efd = efd

	event := &unix.EpollEvent{
		Fd:     int32(e.sfd),
		Events: unix.EPOLLIN,
	}

	return unix.EpollCtl(efd, unix.EPOLL_CTL_ADD, e.sfd, event)
}

func (e *EpollListener) Listen(port int) error {
	sfd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, unix.IPPROTO_TCP)
	if err != nil {
		return err
	}
	e.sfd = sfd

	sa := &unix.SockaddrInet4{}
	sa.Addr = [4]byte{0, 0, 0, 0}
	sa.Port = port
	if err = unix.Bind(sfd, sa); err != nil {
		return err
	}

	if err = unix.Listen(sfd, 1024); err != nil {
		return err
	}
	e.Info("server listen", zap.String("addr", "0.0.0.0"), zap.Int("port", port))

	return nil
}

func (e *EpollListener) Run() error {
	if err := e.epollInit(); err != nil {
		return err
	}
	go e.db.PurgePeriodically()

	events := make([]unix.EpollEvent, 1024)
	for {
		n, err := unix.EpollWait(e.efd, events, 200)
		if err != nil {
			if e, ok := err.(syscall.Errno); ok && int(e) == 4 {
				continue
			}
			e.Warn("epoll_wail fail", zap.Error(err))
		}
		if n <= 0 {
			e.db.DelExpireKey()
			continue
		}
		for i := 0; i < n; i++ {
			event := events[i]
			if event.Fd == int32(e.sfd) {
				cfd, _, err := unix.Accept(e.sfd)
				if err != nil {
					e.Error("accept connection fail", zap.Error(err))
					continue
				}
				ee := &unix.EpollEvent{
					Fd:     int32(cfd),
					Events: unix.EPOLLIN | unix.EPOLLET,
				}
				if err = unix.EpollCtl(e.efd, unix.EPOLL_CTL_ADD, cfd, ee); err != nil {
					e.Warn("epoll_ctl new connection fail", zap.Error(err))
				}
			} else {
				cfd := int(event.Fd)
				if event.Events&unix.EPOLLRDHUP != 0 || event.Events&unix.EPOLLHUP != 0 {
					e.closeClientFd(cfd)
					continue
				}

				if event.Events&unix.EPOLLIN != 0 {
					f, n, err := readFrame(cfd)
					if err != nil {
						e.Error("read frmae fail", zap.Error(err))
						continue
					}
					if n == 0 {
						e.closeClientFd(cfd)
						continue
					}
					if len(f) == 0 {
						unix.Write(cfd, unknownProto)
						continue
					}

					cmd := f.GetCmd()
					replyBytes := cmd(e.db, f.GetData())
					if _, err = unix.Write(cfd, replyBytes); err != nil {
						e.Error("write reply fail", zap.Error(err))
					}
				}
			}
		}
	}
}

func (e *EpollListener) Stop() {
	e.Info("server closing")
	unix.Close(e.efd)
	unix.Close(e.sfd)
	e.db.Close()
	e.Info("server close done")
}

func (e *EpollListener) closeClientFd(cfd int) {
	unix.Shutdown(cfd, unix.SHUT_RDWR)
	unix.Close(cfd)
	unix.EpollCtl(e.efd, unix.EPOLL_CTL_DEL, cfd, nil)
}

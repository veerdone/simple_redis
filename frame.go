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
	"github.com/veerdone/simple_redis/command"
	"github.com/veerdone/simple_redis/util"
	"golang.org/x/sys/unix"
)

var unknownProto = []byte("unknown protocal")

// Frame struct
//
// | frame length | cmd length | cmd | data length | data |
// | 32 bit		  | 4 bit      | ... | 32 bit      | ...  |
type Frame []byte

func (f Frame) GetCmd() command.Cmd {
	cmdBytes := f[1 : 1+f[0]]
	str := string(cmdBytes)
	
	return command.GetCommand(str)
}

func (f Frame) GetData() []byte {
	l := 1 + f[0]
	dataLengthBytes := f[l : l+4]
	dataLength := util.BytesToInt(dataLengthBytes)

	return f[l+4 : int(l+4)+dataLength]
}

func readFrame(fd int) (f Frame, n int, err error) {
	lenBytes := make([]byte, 4)
	n, err = unix.Read(fd, lenBytes)
	if err != nil || n == 0 {
		return
	}

	payloadLen := util.BytesToInt(lenBytes)
	if payloadLen == 0 {
		return
	}

	paylodBytes := make([]byte, payloadLen)
	n, err = unix.Read(fd, paylodBytes)
	if err != nil || n == 0 {
		return
	}
	f = Frame(paylodBytes)

	return
}

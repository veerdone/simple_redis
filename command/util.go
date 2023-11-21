package command

import (
	"bytes"

	"github.com/veerdone/simple_redis/db"
)

type splitData struct {
	splitBytes [][]byte
	errBytes []byte
	key string
	entity *db.Entity
}

func CheckArgsNumAndKeyExist(b []byte, n int, d *db.DB) splitData {
	sd := splitData{}

	sd.splitBytes = bytes.SplitN(b, spaceSplit, n)
	if len(sd.splitBytes) != n {
		sd.errBytes = errResp(WrongNumArgReply)

		return sd
	}
	sd.key = string(sd.splitBytes[0])
	sd.entity = d.Get(sd.key)
	if sd.entity == nil {
		sd.errBytes = errResp(NilReply)
	}

	return sd
}

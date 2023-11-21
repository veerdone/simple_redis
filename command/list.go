package command

import (
	"bytes"
	"strconv"

	"github.com/veerdone/simple_redis/db"
)

func LPushCmd(d *db.DB, data []byte) []byte {
	splitBytes := bytes.SplitN(data, spaceSplit, 2)
	if len(splitBytes) != 2 {
		return errResp(WrongNumArgReply)
	}

	key := string(spaceSplit[0])
	entity := d.Get(key)
	if entity == nil {
		return errResp(NilReply)
	}
	if entity.Types != db.LIST {
		return errResp(WrongTypeReply)
	}
	list := (*db.List)(entity.Data)
	list.LPush(splitBytes[1])

	return strResp(OKReply)
}

func LPopCmd(d *db.DB, data []byte) []byte {
	key := string(data)

	entity := d.Get(key)
	if entity == nil {
		return errResp(NilReply)
	}
	if entity.Types != db.LIST {
		return errResp(WrongTypeReply)
	}

	list := (*db.List)(entity.Data)
	reply := list.LPop()
	if list.Len() == 0 {
		d.Del(key)
	}

	return dataResp(reply)
}

func RPushCmd(d *db.DB, data []byte) []byte {
	splitBytes := bytes.SplitN(data, spaceSplit, 2)
	if len(splitBytes) != 2 {
		return errResp(WrongNumArgReply)
	}

	key := string(spaceSplit[0])
	entity := d.Get(key)
	if entity == nil {
		return errResp(NilReply)
	}
	if entity.Types != db.LIST {
		return errResp(WrongTypeReply)
	}
	list := (*db.List)(entity.Data)
	list.RPush(splitBytes[1])

	return strResp(OKReply)
}

func RPopCmd(d *db.DB, data []byte) []byte {
	key := string(data)

	entity := d.Get(key)
	if entity == nil {
		return errResp(NilReply)
	}
	if entity.Types != db.LIST {
		return errResp(WrongTypeReply)
	}

	list := (*db.List)(entity.Data)
	reply := list.RPop()
	if list.Len() == 0 {
		d.Del(key)
	}

	return dataResp(reply)
}

func LIndexCmd(d *db.DB, data []byte) []byte {
	splitBytes := bytes.SplitN(data, spaceSplit, 2)
	if len(splitBytes) != 2 {
		return errResp(WrongNumArgReply)
	}

	key := string(splitBytes[0])
	entity := d.Get(key)
	if entity == nil {
		return errResp(NilReply)
	}
	if entity.Types != db.LIST {
		return errResp(WrongTypeReply)
	}
	index, err := strconv.Atoi(string(splitBytes[1]))
	if err != nil {
		return errResp(WrongValue)
	}
	list := (*db.List)(entity.Data)

	return dataResp(list.Index(index))
}

func LRangeCmd(d *db.DB, data []byte) []byte {
	splitBytes := bytes.SplitN(data, spaceSplit, 3)
	if len(splitBytes) != 3 {
		return errResp(WrongNumArgReply)
	}

	key := string(splitBytes[0])
	entity := d.Get(key)
	if entity == nil {
		return errResp(NilReply)
	}
	if entity.Types != db.LIST {
		return errResp(WrongTypeReply)
	}
	start, err := strconv.Atoi(string(splitBytes[1]))
	if err != nil {
		return errResp(WrongValue)
	}
	end, err := strconv.Atoi(string(splitBytes[2]))
	if err != nil {
		return errResp(WrongValue)
	}
	list := (*db.List)(entity.Data)

	return multiResp(list.Range(start, end))
}
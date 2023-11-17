package command

import (
	"bytes"
	"strconv"

	"github.com/veerdone/simple_redis/db"
)

func LPushCmd(d *db.DB, data []byte) []byte {
	splitBytes := bytes.SplitN(data, spaceSplit, 2)
	if len(splitBytes) != 2 {
		return WrongNumArgReply
	}

	key := string(spaceSplit[0])
	entity := d.Get(key)
	if entity == nil {
		return NilReply
	}
	if entity.Types != db.LIST {
		return WrongTypeReply
	}
	list := (*db.List)(entity.Data)
	list.LPush(splitBytes[1])

	return OKReply
}

func LPopCmd(d *db.DB, data []byte) []byte {
	key := string(data)

	entity := d.Get(key)
	if entity == nil {
		return NilReply
	}
	if entity.Types != db.LIST {
		return WrongTypeReply
	}

	list := (*db.List)(entity.Data)
	reply := list.LPop()
	if list.Len() == 0 {
		d.Del(key)
	}

	return reply
}

func RPushCmd(d *db.DB, data []byte) []byte {
	splitBytes := bytes.SplitN(data, spaceSplit, 2)
	if len(splitBytes) != 2 {
		return WrongNumArgReply
	}

	key := string(spaceSplit[0])
	entity := d.Get(key)
	if entity == nil {
		return NilReply
	}
	if entity.Types != db.LIST {
		return WrongTypeReply
	}
	list := (*db.List)(entity.Data)
	list.RPush(splitBytes[1])

	return OKReply
}

func RPopCmd(d *db.DB, data []byte) []byte {
	key := string(data)

	entity := d.Get(key)
	if entity == nil {
		return NilReply
	}
	if entity.Types != db.LIST {
		return WrongTypeReply
	}

	list := (*db.List)(entity.Data)
	reply := list.RPop()
	if list.Len() == 0 {
		d.Del(key)
	}

	return reply
}

func LIndexCmd(d *db.DB, data []byte) []byte {
	splitBytes := bytes.SplitN(data, spaceSplit, 2)
	if len(splitBytes) != 2 {
		return WrongNumArgReply
	}

	key := string(spaceSplit[0])
	entity := d.Get(key)
	if entity == nil {
		return NilReply
	}
	if entity.Types != db.LIST {
		return WrongTypeReply
	}
	index, err := strconv.Atoi(string(splitBytes[1]))
	if err != nil {
		return WrongValue
	}
	list := (*db.List)(entity.Data)

	return list.Index(index)
}

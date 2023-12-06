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


package command

import (
	"strconv"

	"github.com/veerdone/simple_redis/db"
)

func LPushCmd(d *db.DB, data []byte) []byte {
	sd := CheckArgsNumAndKeyExist(data, 2, d)
	if len(sd.errBytes) != 0 {
		return sd.errBytes
	}

	entity := sd.entity
	splitBytes := sd.splitBytes
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
	sd := CheckArgsNumAndKeyExist(data, 2, d)
	if len(sd.errBytes) != 0 {
		return sd.errBytes
	}

	entity := sd.entity
	splitBytes := sd.splitBytes
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
	sd := CheckArgsNumAndKeyExist(data, 2, d)
	if len(sd.errBytes) != 0 {
		return sd.errBytes
	}

	entity := sd.entity
	splitBytes := sd.splitBytes
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
	sd := CheckArgsNumAndKeyExist(data, 3, d)
	if len(sd.errBytes) != 0 {
		return sd.errBytes
	}

	entity := sd.entity
	splitBytes := sd.splitBytes
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

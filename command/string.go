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
	"bytes"
	"strconv"
	"unsafe"

	"github.com/veerdone/simple_redis/db"
)

func SetCmd(d *db.DB, data []byte) []byte {
	splitBytes := bytes.SplitN(data, spaceSplit, 2)
	if len(splitBytes) != 2 {
		return WrongNumArgReply
	}
	key := string(splitBytes[0])
	value := splitBytes[1]
	entity := d.Get(key)
	if entity == nil {
		entity = &db.Entity{}
	}
	entity.Data = unsafe.Pointer(&value)
	entity.Types = db.STRING

	v, err := strconv.ParseInt(string(value), 10, 64)
	if err == nil {
		entity.Data = unsafe.Pointer(&v)
		entity.Types = db.INT
	}
	d.Set(key, entity)

	return OKReply
}

func GetCmd(d *db.DB, data []byte) []byte {
	key := string(data)
	entity := d.Get(key)
	if entity == nil {
		return NilReply
	}
	switch entity.Types {
	case db.STRING:
		s := (*[]byte)(entity.Data)
		return *s
	case db.INT:
		i := (*int64)(entity.Data)
		return []byte(strconv.FormatInt(*i, 10))
	}

	return WrongTypeReply
}

func IncrCmd(d *db.DB, data []byte) []byte {
	key := string(data)
	entity := d.Get(key)
	if entity == nil {
		return NilReply
	}
	if entity.Types != db.INT {
		return WrongTypeReply
	}

	num := (*int64)(entity.Data)
	*num = *num + 1

	return []byte(strconv.FormatInt(*num, 10))
}

func DecrCmd(d *db.DB, data []byte) []byte {
	key := string(data)
	entity := d.Get(key)
	if entity == nil {
		return NilReply
	}
	if entity.Types != db.INT {
		return WrongTypeReply
	}

	num := (*int64)(entity.Data)
	*num = *num - 1

	return []byte(strconv.FormatInt(*num, 10))
}

func IncrByCmd(d *db.DB, data []byte) []byte {
	splitBytes := bytes.SplitN(data, spaceSplit, 2)
	if len(splitBytes) != 2 {
		return WrongNumArgReply
	}
	key := string(spaceSplit[0])
	incrNum, err := strconv.ParseInt(string(spaceSplit[1]), 10, 64)
	if err != nil {
		return WrongValue
	}
	entity := d.Get(key)
	if entity == nil {
		return NilReply
	}

	num := (*int64)(entity.Data)
	*num += incrNum

	return []byte(strconv.FormatInt(*num, 10))
}

func DecrByCmd(d *db.DB, data []byte) []byte {
	splitBytes := bytes.SplitN(data, spaceSplit, 2)
	if len(splitBytes) != 2 {
		return WrongNumArgReply
	}
	key := string(spaceSplit[0])
	decrNum, err := strconv.ParseInt(string(spaceSplit[1]), 10, 64)
	if err != nil {
		return WrongValue
	}
	entity := d.Get(key)
	if entity == nil {
		return NilReply
	}

	num := (*int64)(entity.Data)
	*num -= decrNum

	return []byte(strconv.FormatInt(*num, 10))
}
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

import "github.com/veerdone/simple_redis/db"

type Cmd func(db *db.DB, data[] byte) []byte

var (
	OKReply = []byte("OK\r\n")
	UnknownReply = []byte("unknow cmd\r\n")
	NilReply = []byte("null\r\n")
	WrongTypeReply = []byte("WRONGTYPE Operation against a key holding the wrong kind of value\r\n")
	WrongNumArgReply = []byte("ERR wrong number of arguments for command\r\n")
    WrongValue = []byte("Err wrong vlaue's type is not integer")
)

var spaceSplit = []byte{32}

func UnknownCmd(db *db.DB, data[] byte) []byte {
	return UnknownReply
}
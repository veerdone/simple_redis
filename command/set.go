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


func SAdd(d *db.DB, data []byte) []byte {
	sd := CheckArgsNumAndKeyExist(data, 2, d)
	if len(sd.errBytes) != 0 {
		return sd.errBytes
	}
	if sd.entity.Types != db.SET {
        return WrongTypeReply   
    }
	set := (*db.HashSet)(sd.entity.Data)
    set.Add(sd.splitBytes[1])

	return strResp(OKReply)
}

func SIsMember(d *db.DB, data []byte) []byte {
    sd := CheckArgsNumAndKeyExist(data, 2, d)
    if len(sd.errBytes) != 0 {
        return sd.errBytes
    }
    if sd.entity.Types != db.SET {
        return WrongTypeReply
    }
    set := (*db.HashSet)(sd.entity.Data)
    if set.Contains(sd.splitBytes[1]) {
        return numberResp(numberOneBytes)
    }

    return numberResp(numberZeroBytes)
}
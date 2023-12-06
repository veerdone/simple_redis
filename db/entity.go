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

package db

import "unsafe"

const (
	STRING uint8 = 1
	INT    uint8 = 2
	LIST   uint8 = 3
    SET uint8 = 4
)

type Entity struct {
	Types    uint8
	Data     unsafe.Pointer
	LastRead int64
	ExpireAt int64
}

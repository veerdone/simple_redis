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

package util

func BytesToInt(b []byte) int {
	value := 0
	value = int((b[3] & 0xFF) | ((b[2] & 0xFF) << 8) | ((b[1] & 0xFF) << 16) | ((b[0] & 0xFF) << 24))

	return value
}

func IntToBytes(i int, b []byte) []byte {
	if len(b) == 0 {
		b = make([]byte, 4)
	}
	b[0] = byte((i>>24) & 0xFF)
	b[1] = byte((i>>16) & 0xFF)
	b[2] = byte((i>>8) & 0xFF)
	b[3] = byte(i & 0xFF)

	return b
}
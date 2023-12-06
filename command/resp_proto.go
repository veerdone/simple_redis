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

	"github.com/veerdone/simple_redis/util"
)

// response protocol
// | header       | frame                            |
// | frame length | type  | payload length | payload |
// | 4 bit        | 1 bit | 4 bit          | ...     |
const (
	strRespProto      = 1
	errRespProto      = 2
	numberRespProto   = 3
	multiValRespProto = 4
	dataRespProto     = 5
)

var (
	numberOneBytes = []byte("1")
	numberZeroBytes = []byte("0")
	multiValSep = []byte("\r\n")
)

func strResp(b []byte) []byte {
	return buildResp(strRespProto, b)
}

func errResp(b []byte) []byte {
	return buildResp(errRespProto, b)
}

func numberResp(b []byte) []byte {
	return buildResp(numberRespProto, b)
}

func dataResp(b []byte) []byte {
	return buildResp(dataRespProto, b)
}

func multiResp(b [][]byte) []byte {
	data := bytes.Join(b, multiValSep)

	return buildResp(multiValRespProto, data)
}

func buildResp(proto int, data []byte) []byte {
	frame := make([]byte, 0, 5+len(data))
	frame = append(frame, byte(proto))
	frame = append(frame, util.IntToBytes(len(data), nil)...)
	frame = append(frame, data...)

	resp := make([]byte, 0, 4+len(frame))
	resp = append(resp, util.IntToBytes(len(frame), nil)...)
	resp = append(resp, frame...)

	return resp
}

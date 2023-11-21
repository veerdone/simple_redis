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

var multiValSep = []byte("\r\n")

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

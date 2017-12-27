package handler

import (
	"encoding/hex"
	"github.com/gwtony/gapi/log"
)

type UAgentUsocketHandler struct {
	token  string
	log    log.Log
}

func (handler *UAgentUsocketHandler) ServUsocket(data []byte, size int) {
	//TODO: check magic
	id := make([]byte, UNIQID_SIZE/2)
	handler.log.Debug("Deal usocket:", string(data[0:UNIQID_SIZE]))
	ret, err := hex.Decode(id, data[0:UNIQID_SIZE])
	//TODO: error encoding/hex: invalid byte: U+0000
	if err != nil && ret != 0 {
		handler.log.Error("Hex decode error:", err)
		return
	}

	rdata := make([]byte, size)
	copy(rdata, data[:size])

	if size > UAGENT_DEFAULT_TCP_SIZE {
		handler.log.Error("Data size is over limit(%d)", UAGENT_DEFAULT_TCP_SIZE)
		return
	}

	SendToRouter(rdata, size)
}

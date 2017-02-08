package handler

import (
	"fmt"
	"encoding/hex"
	"github.com/gwtony/gapi/log"
)

// AddHandler urouter udp handler
type UAgentUsocketHandler struct {
	token  string
	log    log.Log
}

func (handler *UAgentUsocketHandler) ServUsocket(data []byte, size int) {
	//TODO: check magic
	id := make([]byte, UNIQID_SIZE/2)
	handler.log.Debug("Deal udp:", string(data[0:UNIQID_SIZE]))
	ret, err := hex.Decode(id, data[0:UNIQID_SIZE])
	//TODO: error encoding/hex: invalid byte: U+0000
	if err != nil && ret != 0 {
		fmt.Println(err, ret)
	}

	rdata := make([]byte, size)
	copy(rdata, data[:size])

	if size > 65535 {
		handler.log.Error("Data size is over limit(65535)")
		return
	}

	SendToRouter(rdata, size)
}

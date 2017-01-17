package handler

import (
	//"net"
	"fmt"
	//"time"
	//"math/rand"
	//"strings"
	//"strconv"
	//"io/ioutil"
	//"net/http"
	//"encoding/json"
	//"encoding/binary"
	"encoding/hex"
	//"bytes"
	//"gopkg.in/redis.v5"
	//"github.com/ugorji/go/codec"
	"github.com/gwtony/gapi/log"
	//"github.com/gwtony/gapi/api"
	//"github.com/gwtony/gapi/errors"
)

// AddHandler urouter udp handler
type UAgentUsocketHandler struct {
	token  string
	log    log.Log
}

func (handler *UAgentUsocketHandler) ServUsocket(data []byte, size int) {
	//TODO: check magic

	//var ids string
	id := make([]byte, UNIQID_SIZE/2)
	fmt.Println("in deal udp", string(data[0:UNIQID_SIZE]))
	ret, err := hex.Decode(id, data[0:UNIQID_SIZE])
	//TODO: error encoding/hex: invalid byte: U+0000
	if err != nil && ret != 0 {
		fmt.Println(err, ret)
//	} else {
//		fmt.Printf("ID: ")
//		for _, i := range id {
//			fmt.Printf("%02x", i)
//			ids += fmt.Sprintf("%02X", i)
//		}
//		fmt.Println()
//		fmt.Println(ids)
	}

	//fmt.Println(data[0:32])
	//fmt.Println(string(data[78:size]))
	rdata := make([]byte, size)
	copy(rdata, data[:size])

	if size <= 65535 {
		SendToRouter(rdata, size)
		//SendToRouter(data, size)
	}

		//var id string
//	var rh UDPdata
//	var mh codec.MsgpackHandle

	//id = string(data[0:UNIQID_SIZE])
//	handler.log.Debug("Got id: %s", id)

//	b_buf := bytes.NewBuffer(data[UNIQID_SIZE + 2:size])

//	r := bytes.NewReader(b_buf.Bytes())
//	dec := codec.NewDecoder(r, &mh)
//	err := dec.Decode(&rh)
//	if err != nil {
//		handler.log.Debug(err)
//	} else {
//		handler.log.Debug("{Uid: %s, Puid: %s, Pip: %s, Pport: %d, Lip: %s, Lport: %d, Dlen: %d}",
//				string(rh.Uid), string(rh.Puid), string(rh.Pip), rh.Pport, string(rh.Lip), rh.Lport, rh.Dlen)
//		//handler.log.Debug("Data :", string(rh.Data))
//	}

	//handler.rh.Set(id, data[UNIQID_SIZE + 2:size], UROUTER_DEFAULT_TTL)
}

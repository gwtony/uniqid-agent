package handler

import (
	"net"
	"sync"
	"time"
	"strings"
	"encoding/binary"
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/errors"
)

type TcpHandler struct {
	name  string
	port  string

	lock  *sync.RWMutex
	amap  map[string]int
	addrs []string
	ch    chan *RouterData
	cch   chan int
	log   log.Log
}

var thandler *TcpHandler

func (th *TcpHandler)tcpWorker(ip, port string) {
	var data *RouterData
	var n int

	addr :=  ip + ":" + port
	th.log.Debug("Tcp worker start with %s", addr)

	conn, err := net.DialTimeout("tcp", addr, 3 * time.Second)
	if err != nil {
		th.log.Error("Connect to %s failed: %s", addr, err)
		th.lock.Lock()
		th.amap[ip] = 0
		th.lock.Unlock()
		th.cch<-1
		return
	}

	defer conn.Close()

	for {
		select {
		case data = <-th.ch:
			bdata := make([]byte, data.Len + 3)
			bdata[0] = data.Magic
			th.log.Debug("data.Len is %d, uid is %s", data.Len, data.Data[0:32])
			binary.LittleEndian.PutUint16(bdata[1:3], data.Len)
			copy(bdata[3:], data.Data)

			err = conn.SetWriteDeadline(time.Now().Add(time.Second))
			if err != nil {
				th.log.Error("Set write deadline to %s failed: %s", addr, err)
				th.lock.Lock()
				th.amap[ip] = 0
				th.lock.Unlock()
				th.cch<-1
				return
			}

			n, err = conn.Write(bdata)
			if err != nil {
				//TODO: connection failed, retry
				th.log.Error("Write to %s failed: %s", addr, err)
				//TODO: if connection is broken
				th.lock.Lock()
				th.amap[ip] = 0
				th.lock.Unlock()
				th.cch<-1
				return
			}
			th.log.Debug("Write %d data to %s", n, addr)
		}
	}

	th.log.Debug("quit worker")
}

func (th *TcpHandler)tcpMonitor() {
	for {
		select {
		case <- th.cch:
			th.log.Debug("Got quit")
			time.Sleep(time.Second)
			ns, err := net.LookupHost(th.name)
			if err != nil {
				th.log.Error("Look up host %s failed", th.name)
				break
			}
			if len(ns) < 1 {
				th.log.Error("Look up host %s no avaliable addrs", th.name)
				break
			}
			for _, ip := range ns {
				th.lock.Lock()
				if th.amap[ip] == 0 {
					th.amap[ip] = 1
					go th.tcpWorker(ip, th.port)
				}
				th.lock.Unlock()
			}
			th.addrs = ns
		}
	}
}

func InitTcpHandler(addr string, log log.Log) error {
	thandler = &TcpHandler{
		amap: make(map[string]int),
		lock: &sync.RWMutex{},
		ch  : make(chan *RouterData, 1024),
		cch : make(chan int, 10),
		log : log,
	}

	name := strings.Split(addr, ":")
	if len(name) != 2 {
		return InitTcpHandlerError
	}
	thandler.name = name[0]
	thandler.port = name[1]

	ns, err := net.LookupHost(name[0])
	if err != nil {
		log.Error("Look up host %s failed", name[0])
		return errors.LookupHostError
	}
	if len(ns) < 1 {
		log.Error("Look up host %s failed, no avaliable addrs", name[0])
		return errors.LookupHostError
	}

	thandler.addrs = ns
	//thandler.ch = make(chan RouterData, 1024)
	//thandler.cch = make(chan int, 10)

	for _, ip := range thandler.addrs {
		thandler.lock.Lock()
		thandler.amap[ip] = 1
		thandler.lock.Unlock()
		go thandler.tcpWorker(ip, thandler.port)
	}

	go thandler.tcpMonitor()
	return nil
}

func SendToRouter(data []byte, size int) error {
	rdata := &RouterData{}
	rdata.Magic = byte(0x77)
	rdata.Len = uint16(size)
	rdata.Data = data
	thandler.log.Debug("Send to router: data len is %d, uid is %s", rdata.Len, data[0:32])
	thandler.ch <- rdata
	return nil
}

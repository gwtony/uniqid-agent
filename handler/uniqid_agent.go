package handler

import (
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/gapi/config"
)

// InitContext inits uniqid agent context
func InitContext(conf *config.Config, log log.Log) error {
	cf := &UAgentConfig{}
	err := cf.ParseConfig(conf)
	if err != nil {
		log.Error("Uniqid agent parse config failed")
		return err
	}

	err = InitTcpHandler(cf.raddr, log)
	if err != nil {
		log.Error("Uniqid agent init tcp handler failed")
		return err
	}

	//apiLoc := cf.apiLoc
	//token  := cf.token


	api.AddUsocketHandler(&UAgentUsocketHandler{log: log})
	//api.AddHttpHandler(apiLoc + MACEDON_ADD_LOC, &AddHandler{h: h, domain: domain, pc: pc, token: token, log: log})
	//api.AddHttpHandler(apiLoc + MACEDON_DELETE_LOC, &DeleteHandler{h: h, domain: domain, pc: pc, token: token, log: log})
	//api.AddHttpHandler(apiLoc + MACEDON_READ_LOC, &ReadHandler{h: h, domain: domain, token: token, log: log})
	//api.AddHttpHandler(apiLoc + MACEDON_SCAN_LOC, &ScanHandler{h: h, domain: domain, token: token, log: log})
	//api.AddHttpHandler(apiLoc + MACEDON_ADD_SERVER_LOC, &AddServerHandler{h: h, pc: pc, token: token, log: log})
	//api.AddHttpHandler(apiLoc + MACEDON_DELETE_SERVER_LOC, &DeleteServerHandler{h: h, pc: pc, token: token, log: log})
	//api.AddHttpHandler(apiLoc + MACEDON_READ_SERVER_LOC, &ReadServerHandler{h: h, token: token, log: log})

	return nil
}



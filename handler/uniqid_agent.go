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

	api.AddUsocketHandler(&UAgentUsocketHandler{log: log})

	return nil
}



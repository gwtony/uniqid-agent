package main

import (
	"fmt"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/uniqid_agent/handler"
)

func main() {
	err := api.Init()
	if err != nil {
		fmt.Println("Init api failed")
		return
	}
	api.SetConfig("uniqid_agent.conf")
	config := api.GetConfig()
	log := api.GetLog()

	err = handler.InitContext(config, log)
	if err != nil {
		fmt.Println("Init uniqid agent failed")
		return
	}

	api.Run()
}

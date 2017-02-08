package handler

import (
	"fmt"
	"os"
	"time"
	"github.com/gwtony/gapi/config"
	"github.com/gwtony/gapi/errors"
)

// URouterConfig URouter config
type UAgentConfig struct {
	raddr      string /* uniqid router addr */

	timeout    time.Duration
}

// ParseConfig parses config
func (conf *UAgentConfig) ParseConfig(cf *config.Config) error {
	var err error
	if cf.C == nil {
		return errors.BadConfigError
	}
	raddrStr, err := cf.C.GetString("uagent", "urouter_addr")
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] [uagent] Read conf: No urouter_addr")
		return err
	}
	if raddrStr == "" {
		fmt.Fprintln(os.Stderr, "[Error] [uagent] Empty urouter server address")
		return errors.BadConfigError
	}
	conf.raddr = raddrStr
	//raddr := strings.Split(raddrStr, ",")
	//for i := 0; i < len(raddr); i++ {
	//	if raddr[i] != "" {
	//		if !strings.Contains(raddr[i], ":") {
	//			conf.raddr = append(conf.raddr, raddr[i] + ":" + DEFAULT_REDIS_PORT)
	//		} else {
	//			conf.raddr = append(conf.raddr, raddr[i])
	//		}
	//	}
	//}

	timeout, err := cf.C.GetInt64("uagent", "timeout")
	if err != nil || timeout <= 0 {
		fmt.Fprintln(os.Stderr, "[Info] [uagent] Read conf: use default timeout: ", UAGENT_DEFAULT_TIMEOUT)
		timeout = UAGENT_DEFAULT_TIMEOUT
	}
	conf.timeout =  time.Duration(timeout) * time.Second

	return nil
}

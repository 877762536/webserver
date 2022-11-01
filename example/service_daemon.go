package example

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"webserver"
	"webserver/config"
	"webserver/svc"
)

const _groupService = "service"

// 不支持windows系统

type ServiceConfig struct {
	PidFile   string `json:"pid_file"`
	DaemonLog string `json:"daemon_log"`
}

func getConf() (ServiceConfig, error) {
	var c ServiceConfig
	err := config.GetConfig(_groupService, &c)
	if err != nil {
		return c, err

	}
	confWebServer, err := webserver.GetConfig()
	if err != nil {
		return c, err
	}
	logPath := confWebServer.LogPath
	c.PidFile = filepath.Join(logPath, c.PidFile)
	c.DaemonLog = filepath.Join(logPath, c.DaemonLog)

	if err := os.MkdirAll(logPath, 0755); err != nil {
		return c, fmt.Errorf("failed to create log dir %s", logPath)
	}

	return c, nil
}

func init() {
	c := ServiceConfig{
		PidFile:   "console.pid",
		DaemonLog: "daemon.log",
	}
	_ = config.RegConfig(_groupService, c)
}

func NewService(prc *MainProc) (*svc.Svc, error) {
	//if runtime.GOOS == "windows" {
	//	svcName := "iguard6center"
	//	svcDescription := "iguard6 center 服务"
	//	return svc.NewSvc(svcName, svcDescription, prc), nil
	//} else {
	//	return svc.NewSvc(c.PidFile, c.DaemonLog, prc), nil
	//}

	c, err := getConf()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to get daemon config:%v", err))
	}
	return svc.NewSvc(c.PidFile, c.DaemonLog, prc), nil
}

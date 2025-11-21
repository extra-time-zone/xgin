package xgin

import (
	"github.com/extra-time-zone/xgin/logger"
	"github.com/extra-time-zone/xgin/xglobal"
)

func SetConfig(conf *Config) {
	if conf.GinMode != "" {
		xglobal.GinMode = conf.GinMode
	}
	if conf.LogFile != "" {
		xglobal.LogFile = conf.LogFile
	}

	//init zero log
	if xglobal.LogFile != "" {
		logger.Init(xglobal.LogFile)
	} else {
		logger.Init()
	}
}

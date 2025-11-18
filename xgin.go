package xgin

import (
	"fmt"

	"github.com/extra-time-zone/xgin/logger"
	"github.com/extra-time-zone/xgin/xglobal"
)

func SetConfig(conf *Config) {
	if conf.GinMode != "" {
		xglobal.GinMode = conf.GinMode
	}
	fmt.Printf("---conf: %+v\n", conf)
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

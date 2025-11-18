package iface

import (
	"github.com/extra-time-zone/xgin/request"
	"github.com/extra-time-zone/xgin/response"
	"github.com/extra-time-zone/xgin/xerror"
)

type IHandler interface {
	PreHandler(*request.Request)
	Handler(*request.Request) (*response.Response, xerror.Error)
	PostHandler(*request.Request)
}

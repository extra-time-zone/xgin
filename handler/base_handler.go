package handler

import (
	"github.com/extra-time-zone/xgin/request"
	"github.com/extra-time-zone/xgin/response"
	"github.com/extra-time-zone/xgin/xerror"
)

type BaseHandle struct{}

func (c *BaseHandle) PreHandler(*request.Request) {}
func (c *BaseHandle) Handler(*request.Request) (*response.Response, *xerror.XError) {
	return &response.Response{}, nil
}
func (c *BaseHandle) PostHandler(*request.Request) {}

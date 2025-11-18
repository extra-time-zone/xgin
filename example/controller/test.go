package controller

import (
	"fmt"
	"time"

	"github.com/extra-time-zone/xgin/example/service"
	"github.com/extra-time-zone/xgin/handler"
	"github.com/extra-time-zone/xgin/request"
	"github.com/extra-time-zone/xgin/response"
	"github.com/extra-time-zone/xgin/xerror"
	"github.com/spf13/cast"
)

type TestHandler struct {
	handler.BaseHandle
}

func (c *TestHandler) Handler(req *request.Request) (*response.Response, xerror.Error) {
	fmt.Println("TestHandler start...", time.Now().Format(time.DateTime))

	uid := req.GetUid()
	fmt.Println("------uid:", uid)

	x := req.GetCtx().Query("x")
	ix := cast.ToInt(x)

	testService := service.NewTestService(req)
	data, xerr := testService.GetData(ix)
	if xerr != nil {
		return nil, xerror.Wrap(xerr, "test-handler-error")
	}

	return &response.Response{
		Data: data,
	}, nil
}

type TestNotHandler struct {
	handler.BaseHandle
}

func (c *TestNotHandler) Handler(req *request.Request) (*response.Response, xerror.Error) {
	fmt.Println("TestHandler start...", time.Now().Format(time.DateTime))

	testService := service.NewTestService(req)
	data, xerr := testService.GetData(0)
	if xerr != nil {
		return nil, xerror.Wrap(xerr, "test-handler-error")
	}

	return &response.Response{
		Data: data,
	}, nil
}

type ABCHandler struct {
	handler.BaseHandle
}

func (c *ABCHandler) Handler(req *request.Request) (*response.Response, xerror.Error) {
	fmt.Println("TestHandler start...", time.Now().Format(time.DateTime))

	return &response.Response{
		Data: "ABCHandler",
	}, nil
}

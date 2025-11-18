package wrapper

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/extra-time-zone/xgin/iface"
	"github.com/extra-time-zone/xgin/logger"
	"github.com/extra-time-zone/xgin/request"
	"github.com/extra-time-zone/xgin/xerror"
	"github.com/gin-gonic/gin"
)

func HandlerFuncWrapper(handler iface.IHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(c, fmt.Sprintf(`[wrapper panic] XError(start): %v`, err))
				for i := 1; i < 20; i++ {
					if pc, file, line, ok := runtime.Caller(i); ok {
						logger.Error(c, fmt.Sprintf(`[wrapper panic] goroutine:%v, file:%s, function name:%s, line:%d`, pc, file, runtime.FuncForPC(pc).Name(), line))
					}
				}
				logger.Error(c, fmt.Sprintf(`[wrapper panic] XError(end): %v`, err))
			}
		}()

		//handle request
		req := request.NewRequest(c)
		handler.PreHandler(req)
		resp, err := handler.Handler(req)
		handler.PostHandler(req)

		if err != nil {
			//log
			for e := err; e != nil; e = errors.Unwrap(e) {
				logger.Error(req, e.Error())
			}

			//最外层error须为xerror.XError
			if xe := xerror.FirstXError(err); xe != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    req.GetTraceID(),
					"data":    nil,
					"message": xe.Msg,
				})
			}
		} else if resp != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    "0",
				"data":    resp.Data,
				"message": "success",
			})
		}
	}
}

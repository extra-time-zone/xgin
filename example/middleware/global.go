package middleware

import (
	"net/http"

	"github.com/extra-time-zone/xgin/example/pkg/utils"
	"github.com/extra-time-zone/xgin/xglobal"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Global 中间件: 全局中间件
func Global() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.Request.Header.Get(xglobal.TraceId)
		if traceId == "" {
			traceId = uuid.New().String()
		}
		c.Set(xglobal.TraceId, traceId)
		c.Set(xglobal.UserId, "222333")

		//检查IP
		if !verifyIp(c) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
			return
		}

		c.Next()
	}
}

// 检测IP
func verifyIp(c *gin.Context) bool {
	ip := c.ClientIP()

	//本机回路IP
	if utils.IsRequestFromLocal(ip) {
		return true
	}

	return true
}

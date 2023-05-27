package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func init() {
	logger, _ := zap.NewDevelopment()

	zap.ReplaceGlobals(logger)
}
func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		method := c.Request.Method
		uri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		zap.S().Infof("| %3d | %13v | %15s | %s | %s", statusCode, duration, clientIP, method, uri) // 日志格式
	}
}

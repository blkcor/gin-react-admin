package middleware

import (
	"fmt"
	"github.com/blkcor/gin-react-admin/core/logger"
	"github.com/gin-gonic/gin"
)

// RequestLogMiddleware 用于在每次请求时记录请求信息
func RequestLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log requests details
		reqInfo := fmt.Sprintf("%s | %s | %s %s", c.Request.Method, c.Request.URL.Path, c.ClientIP(), c.Request.UserAgent())
		// Process requests
		c.Next()
		//响应信息
		respInfo := fmt.Sprintf("[%d]", c.Writer.Status())
		info := fmt.Sprintf("%s -> %s", reqInfo, respInfo)
		if c.Writer.Status() == 200 {
			logger.Info(info)
		} else {
			logger.Error(info)
		}
	}
}

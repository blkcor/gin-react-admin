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
		info := fmt.Sprintf("%s | %s | %s %s", c.Request.Method, c.Request.URL.Path, c.ClientIP(), c.Request.UserAgent())
		logger.Info(info)
		// Process requests
		c.Next()
	}
}

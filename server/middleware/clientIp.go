package middleware

import (
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

func ClientIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		// 尝试从 X-Forwarded-For 头部获取
		if ip := c.GetHeader("X-Forwarded-For"); ip != "" {
			clientIP = strings.TrimSpace(strings.Split(ip, ",")[0])
		}
		// 尝试从 X-Real-IP 头部获取
		if ip := c.GetHeader("X-Real-IP"); ip != "" {
			clientIP = ip
		}

		// 验证 IP 地址的有效性
		if parsedIP := net.ParseIP(clientIP); parsedIP == nil {
			clientIP = c.RemoteIP()
		}

		// 将解析出的 IP 存储在上下文中
		c.Set("ClientIP", clientIP)
		c.Next()
	}
}

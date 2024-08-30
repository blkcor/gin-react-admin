package middleware

import (
	"github.com/blkcor/gin-react-admin/models/response"
	"github.com/blkcor/gin-react-admin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// OperationLog 操作日志记录中间件
func OperationLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		operator, _ := context.Get("operator")
		context.Next()
		if !strings.Contains(context.Request.URL.Path, "operationLog") {
			err := service.CreateOperationLog(operator.(string), context.Writer.Status(), context.ClientIP(), context.Request.Method, context.Request.URL.Path)
			if err != nil {
				context.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
					Success: false,
					Message: "操作日志记录失败",
					Data:    nil,
				})
				context.Abort()
				return
			}
		}
	}
}

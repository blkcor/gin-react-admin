package middleware

import (
	"github.com/blkcor/gin-react-admin/models/response"
	"github.com/blkcor/gin-react-admin/system"
	"github.com/blkcor/gin-react-admin/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取请求path
		obj := context.Request.URL.Path
		//获取请求方法
		act := context.Request.Method
		//从token中解析处角色
		claim, err := jwt.GetClaimFromContext(context)
		if err != nil {
			context.JSON(http.StatusForbidden, response.BaseResponse[any]{
				Success: false,
				Message: "无权限",
				Data:    nil,
			})
			context.Abort()
			return
		}
		//将username存入上下文，方便操作日志记录
		context.Set("operator", claim.Username)
		sub := claim.RoleCode
		//引入casbin
		e := system.CasbinServiceApp.Casbin()
		//判断策略是否存在
		success, _ := e.Enforce(sub, obj, act)
		//如果环境变量是开发者模式或者casbin校验通过
		if success {
			context.Next()
		} else {
			context.JSON(http.StatusForbidden, response.BaseResponse[any]{
				Success: false,
				Message: "无权限",
				Data:    nil,
			})
			context.Abort()
			return
		}
	}
}

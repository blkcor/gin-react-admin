package middleware

import (
	"github.com/blkcor/gin-react-admin/config/section"
	"github.com/blkcor/gin-react-admin/enums"
	"github.com/blkcor/gin-react-admin/system"
	"github.com/blkcor/gin-react-admin/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !IsUserLogin(context) {
			context.JSON(http.StatusForbidden, gin.H{
				"code":    enums.ToLoginCode,
				"message": "Please login first!",
			})
			context.Abort()
			return
		}
		//user have login
		context.Next()
	}
}

func IsUserLogin(context *gin.Context) bool {
	//get jwt token from request header
	token := context.GetHeader("Authorization")
	//verify if the token is valid
	_, err := jwt.ParseToken(token, section.AppConfig.AccessKey)
	if err != nil {
		return false
	}
	return true
}

func CasbinHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取请求path
		obj := context.Request.URL.Path
		//获取请求方法
		act := context.Request.Method
		//角色应该从token解析出来，此处为了节约时间，写死了值
		sub := "admin"
		//引入casbin
		e := system.CasbinServiceApp.Casbin()
		//判断策略是否存在
		success, _ := e.Enforce(sub, obj, act)
		//如果环境变量是开发者模式或者casbin校验通过
		if success {
			context.Next()
		} else {
			context.JSON(http.StatusForbidden, gin.H{
				"message": "无权限",
			})
			context.Abort()
			return
		}
	}
}

package router

import (
	"github.com/blkcor/gin-react-admin/api"
	"github.com/blkcor/gin-react-admin/core/logger"
	"github.com/blkcor/gin-react-admin/middleware"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	Router = gin.New()
	Router.Use(middleware.CorsMiddleware())
	Router.Use(middleware.RequestLogMiddleware())
	Router.Use(middleware.ClientIP())
	Router.Use(gin.Recovery())
	Router.GET("/captcha", api.Captcha)
	Router.POST("/login", api.Login)
	v1 := Router.Group("/api/v1")
	{
		v1.Use(middleware.AuthMiddleWare())
		v1.Use(middleware.CasbinHandler())
	}
	logger.Info("路由初始化成功!")
}

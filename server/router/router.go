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
	Router.Use(gin.Recovery())
	//session
	Router.GET("/ping", api.Ping)
	//v1 := Router.Group("/api/v1")
	//{
	//	v1.Use(middleware.AuthMiddleWare())
	//	v1.GET("/ping", controller.PingController)
	//}
	logger.Info("路由初始化成功!")
}

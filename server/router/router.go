package router

import (
	"github.com/blkcor/gin-react-admin/api"
	"github.com/blkcor/gin-react-admin/core/logger"
	"github.com/blkcor/gin-react-admin/docs"
	"github.com/blkcor/gin-react-admin/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Router *gin.Engine

func InitDocInfo() {
	docs.SwaggerInfo.Title = "GRA文档"
	docs.SwaggerInfo.Description = "Gin + React 管理系统的 API 文档"
	docs.SwaggerInfo.Version = "0.1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8000"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func Init() {
	InitDocInfo()
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
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.Info("路由初始化成功!")
}

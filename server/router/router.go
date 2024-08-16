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

// @title           GRA API文档
// @version         1.0
// @description     Gin + React 管理系统的 API 文档
// @termsOfService  http://www.swagger.io/terms/

// @contact.name   blkcor
// @contact.url    https://blkcor.me
// @contact.email  blkcor.dev@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Init() {
	//InitDocInfo()
	Router = gin.New()
	Router.Use(middleware.CorsMiddleware())
	Router.Use(middleware.RequestLogMiddleware())
	Router.Use(middleware.ClientIP())
	Router.Use(gin.Recovery())
	Router.GET("/captcha", api.Captcha)
	Router.POST("/login", api.Login)
	vs1 := Router.Group("/v1")
	vs1.Use(middleware.AuthMiddleWare())
	vs1.Use(middleware.CasbinHandler())
	{
		//受保护的路由
	}
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.Info("路由初始化成功!")
}

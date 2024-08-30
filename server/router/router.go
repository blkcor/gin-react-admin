package router

import (
	"github.com/blkcor/gin-react-admin/api"
	v1 "github.com/blkcor/gin-react-admin/api/v1"
	"github.com/blkcor/gin-react-admin/api/ws/system"
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
	docs.SwaggerInfo.Description = "Gin + React 后台管理系统的 API 文档"
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

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Init() {
	//InitDocInfo()
	Router = gin.New()
	Router.Use(middleware.ClientIP())
	Router.Use(middleware.CorsMiddleware())
	Router.Use(middleware.RequestLogMiddleware())
	Router.Use(gin.Recovery())
	Router.GET("/captcha", api.Captcha)
	Router.POST("/login", api.Login)
	Router.POST("/logout", api.Logout)
	protected := Router.Group("/v1")
	protected.Use(middleware.AuthMiddleWare())
	protected.Use(middleware.CasbinHandler())
	protected.Use(middleware.OperationLog())
	{
		//菜单相关接口
		menuGroup := protected.Group("/menu")
		{
			menuGroup.GET("/", v1.GetMenu)
		}

		//操作日志相关接口
		operationLogGroup := protected.Group("/operationLog")
		{
			operationLogGroup.DELETE("/:id", v1.DeleteOperationLogRecord)
			operationLogGroup.DELETE("/deleteOperationLogByIds", v1.DeleteOperationLogByIds)
			operationLogGroup.GET("/", v1.GetOperationLogList)
		}

		resourceMonitor := system.NewResourceMonitor()
		protected.GET("/ws/server-monitor", resourceMonitor.ServerResourceMonitor)
	}

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.Info("路由初始化成功!")
}

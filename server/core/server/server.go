package server

import (
	"fmt"
	"github.com/blkcor/gin-react-admin/config/section"
	"github.com/blkcor/gin-react-admin/core/logger"
	"github.com/blkcor/gin-react-admin/router"
	"github.com/gin-gonic/gin"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)
	//初始化路由
	router.Init()

	logger.Infof("服务启动成功, 地址: http://127.0.0.1:%s", section.AppConfig.Port)
	//开启服务
	fmt.Printf(`
欢迎使用 gin-react-admin
当前版本:v0.0.1
微信号：Collapsar-blkcor
项目地址：https://github.com/blkcor/gin-react-admin
默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
默认前端文件运行地址:http://127.0.0.1:3000
`, section.AppConfig.Port)
	err := router.Router.Run(":" + section.AppConfig.Port)
	if err != nil {
		logger.Errorf("服务启动失败: %s", err.Error())
	}
}

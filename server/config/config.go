package config

import (
	"github.com/blkcor/gin-react-admin/config/section"
	"github.com/blkcor/gin-react-admin/core/logger"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

func Init(done chan bool) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "config.ini")
	conf, err := ini.Load(configPath)
	if err != nil {
		logger.Errorf("Fail to read file: %v", err)
		os.Exit(1)
	}
	app := conf.Section("APP")
	database := conf.Section("POSTGRESQL")
	redis := conf.Section("REDIS")

	section.InitAPP(app)
	section.InitDataBase(database)
	section.InitRedis(redis)

	logger.Info("配置信息初始化完毕!")
	done <- true
}

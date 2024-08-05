package db

import (
	"fmt"
	"github.com/blkcor/gin-react-admin/config/section"
	"github.com/blkcor/gin-react-admin/core/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Init() {
	var err error
	//postgres://username:password@host:port/database_name?options
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", section.DBConfig.User, section.DBConfig.Password, section.DBConfig.Host, section.DBConfig.Port, section.DBConfig.DBName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("fail to init database: ", err)
		os.Exit(1)
	}
	logger.Info("数据库连接已建立!")

	//migration
	//err = DB.AutoMigrate(&model.User{})
	//if err != nil {
	//	logger.Error("🚫 User table migration failed: ", err)
	//	panic(err)
	//}
}

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

	////table migrations
	//_ = DB.AutoMigrate(model.User{})
	//_ = DB.AutoMigrate(model.Role{})
	//_ = DB.AutoMigrate(model.UserRole{})
	//_ = DB.AutoMigrate(model.Menu{})
	//_ = DB.AutoMigrate(model.RoleMenu{})
	//	//创建测试用户
	//	err = createAdmin()
	//	if err != nil {
	//		logger.Error("fail to create admin: ", err)
	//		os.Exit(1)
	//	}
	//	logger.Info("成功创建管理员用户!")
}

//// createAdmin: 创建测试用户
//func createAdmin() (err error) {
//	//user
//	user := model.User{
//		ID:       uuid.New().ID(),
//		Username: "admin",
//		Email:    "blkcor.dev@gmail.com",
//		Avatar:   randx.GetRandomDefaultAvatar(),
//	}
//	err = user.SetPassword("12345678")
//	if err != nil {
//		return
//	}
//	//role
//	role := model.Role{
//		ID:          uuid.New().ID(),
//		RoleName:    "超级管理员",
//		RoleCode:    "admin",
//		Description: "超级管理员角色，拥有最高的权限",
//	}
//	//user_role
//	userRole := model.UserRole{
//		ID:     uuid.New().ID(),
//		UserID: user.ID,
//		RoleID: role.ID,
//	}
//	DB.Create(&user)
//	DB.Create(&role)
//	DB.Create(&userRole)
//	return
//}

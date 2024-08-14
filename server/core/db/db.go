package db

import (
	"fmt"
	"github.com/blkcor/gin-react-admin/config/section"
	"github.com/blkcor/gin-react-admin/core/logger"
	"github.com/blkcor/gin-react-admin/models/model"
	"github.com/google/uuid"
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

	//创建系统菜单
	//createSystemMenu()

}
func createSystemMenu() {
	system := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "系统管理",
		Icon:     "setting",
		Layer:    1,
		ParentID: 0,
	}
	systemUser := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "用户管理",
		Icon:     "user",
		Layer:    2,
		ParentID: system.ID,
	}
	systemRole := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "角色管理",
		Icon:     "grid",
		Layer:    2,
		ParentID: system.ID,
	}
	systemMenu := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "菜单管理",
		Icon:     "fold",
		Layer:    2,
		ParentID: system.ID,
	}
	systemJob := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "任务管理",
		Icon:     "coin",
		Layer:    2,
		ParentID: system.ID,
	}
	monitor := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "系统监控",
		Icon:     "monitor",
		Layer:    1,
		ParentID: 0,
	}
	liveUserPannel := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "在线用户面板",
		Icon:     "collection-tag",
		Layer:    2,
		ParentID: monitor.ID,
	}
	operationLog := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "操作日志",
		Icon:     "operation",
		Layer:    2,
		ParentID: monitor.ID,
	}
	exceptionLog := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "异常日志",
		Icon:     "expand",
		Layer:    2,
		ParentID: monitor.ID,
	}
	serverMonitorPanel := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "服务器监控面板",
		Icon:     "guide",
		Layer:    2,
		ParentID: monitor.ID,
	}
	systemTool := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "系统工具",
		Icon:     "tools",
		Layer:    1,
		ParentID: 0,
	}
	codeGenerator := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "代码生成器",
		Icon:     "help",
		Layer:    2,
		ParentID: systemTool.ID,
	}
	emailTool := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "邮件工具",
		Icon:     "finished",
		Layer:    2,
		ParentID: systemTool.ID,
	}
	mediaLib := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "媒体库",
		Icon:     "set-up",
		Layer:    2,
		ParentID: systemTool.ID,
	}
	imageUpload := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "图片上传",
		Icon:     "upload-filled",
		Layer:    2,
		ParentID: systemTool.ID,
	}
	component := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "组件管理",
		Icon:     "compass",
		Layer:    1,
		ParentID: 0,
	}
	icon := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "图标Icon",
		Icon:     "menu",
		Layer:    2,
		ParentID: component.ID,
	}
	basicComponent := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "基础组件",
		Icon:     "sugar",
		Layer:    2,
		ParentID: component.ID,
	}
	richTextEditor := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "富文本编辑器",
		Icon:     "memo",
		Layer:    2,
		ParentID: component.ID,
	}
	markdownEditor := model.Menu{
		ID:       uuid.New().ID(),
		Name:     "md编辑器",
		Icon:     "edit-pen",
		Layer:    2,
		ParentID: component.ID,
	}
	//插入
	DB.Create(&system)
	DB.Create(&systemUser)
	DB.Create(&systemRole)
	DB.Create(&systemMenu)
	DB.Create(&systemJob)
	DB.Create(&monitor)
	DB.Create(&liveUserPannel)
	DB.Create(&operationLog)
	DB.Create(&exceptionLog)
	DB.Create(&serverMonitorPanel)
	DB.Create(&systemTool)
	DB.Create(&codeGenerator)
	DB.Create(&emailTool)
	DB.Create(&mediaLib)
	DB.Create(&imageUpload)
	DB.Create(&component)
	DB.Create(&icon)
	DB.Create(&basicComponent)
	DB.Create(&richTextEditor)
	DB.Create(&markdownEditor)
	//插入role_menu
	roleMenus := []model.RoleMenu{
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: system.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: systemUser.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: systemRole.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: systemMenu.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: systemJob.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: monitor.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: liveUserPannel.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: operationLog.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: exceptionLog.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: serverMonitorPanel.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: systemTool.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: codeGenerator.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: emailTool.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: mediaLib.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: imageUpload.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: component.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: icon.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: basicComponent.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: richTextEditor.ID},
		{ID: uuid.New().ID(), RoleID: 1169628969, MenuID: markdownEditor.ID},
	}
	DB.Create(&roleMenus)
	logger.Info("成功创建系统菜单!")
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

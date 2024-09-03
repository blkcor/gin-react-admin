package service

import (
	"github.com/blkcor/gin-react-admin/core/db"
	"github.com/blkcor/gin-react-admin/models/model"
)

// GetUserByUsername 根据用户名查找用户
func GetUserByUsername(username string) (model.User, error) {
	var user model.User
	result := db.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

// GetUserRoleByUserId 根据用户id查找对应的角色信息
func GetUserRoleByUserId(userId uint32) (model.Role, error) {
	var userRole model.UserRole
	result := db.DB.Where("user_id = ?", userId).First(&userRole)
	if result.Error != nil {
		return model.Role{}, result.Error
	}
	var role model.Role
	result = db.DB.Where("id = ?", userRole.RoleID).First(&role)
	if result.Error != nil {
		return model.Role{}, result.Error
	}
	return role, nil
}

// GetUserByIds 根据用户id批量查找用户
func GetUserByIds(ids []uint32) ([]model.User, error) {
	var users []model.User
	result := db.DB.Where("id in (?)", ids).Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

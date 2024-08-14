package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primaryKey;comment:用户id"`
	Username  string    `gorm:"type:varchar(100);uniqueIndex;not null;comment:用户名"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null;comment:邮箱"`
	Avatar    string    `gorm:"type:varchar(255);comment:头像"`
	Password  string    `gorm:"type:varchar(255);not null;comment:密码(hashed)"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

// SetPassword 设置hash过后的密码
func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// CheckPassword 检查密码是否正确
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

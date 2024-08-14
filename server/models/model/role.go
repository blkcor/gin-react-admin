package model

import "time"

type Role struct {
	ID          uint32    `gorm:"primaryKey;comment:角色id"`
	RoleName    string    `gorm:"type:varchar(100);uniqueIndex;not null;comment:角色名称"`
	RoleCode    string    `gorm:"type:varchar(50);uniqueIndex;not null;comment:角色编码(admin,member,etc...)"`
	Description string    `gorm:"type:text;comment:角色描述"`
	CreatedAt   time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

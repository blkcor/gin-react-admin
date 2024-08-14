package model

import "time"

type RoleMenu struct {
	ID        uint32    `gorm:"primaryKey;comment:主键"`
	RoleID    uint32    `gorm:"not null;index;uniqueIndex:idx_role_menu;comment:角色id"`
	MenuID    uint32    `gorm:"not null;index;uniqueIndex:idx_role_menu;comment:菜单id"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间"`

	Role Role `gorm:"foreignKey:RoleID"`
	Menu Menu `gorm:"foreignKey:MenuID"`
}

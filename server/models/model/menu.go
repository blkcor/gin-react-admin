package model

import (
	"time"
)

type Menu struct {
	ID        uint32    `gorm:"primaryKey;comment:菜单id"`
	Name      string    `gorm:"type:varchar(100);not null;comment:菜单名称"`
	Icon      string    `gorm:"type:varchar(50);comment:图标"`
	Layer     int       `gorm:"type:int;not null;comment:菜单层级"`
	ParentID  uint32    `gorm:"not null;default:0;comment:父级菜单id"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

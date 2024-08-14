package model

import "time"

type UserRole struct {
	ID        uint32    `gorm:"primaryKey;comment:主键"`
	UserID    uint32    `gorm:"not null;uniqueIndex:idx_user_role;comment:用户id"`
	RoleID    uint32    `gorm:"not null;uniqueIndex:idx_user_role;comment:角色id"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间"`

	User User `gorm:"foreignKey:UserID"`
	Role Role `gorm:"foreignKey:RoleID"`
}

package model

import "time"

type OperationLog struct {
	ID            uint32    `json:"id" gorm:"primary_key"`
	Operator      string    `gorm:"type:varchar(100);not null;comment:操作人"`
	OperateTime   time.Time `gorm:"autoCreateTime;comment:操作时间"`
	StatusCode    int       `gorm:"type:int;not null;comment:状态码"`
	RequestIP     string    `gorm:"type:varchar(100);not null;comment:请求ip"`
	RequestMethod string    `gorm:"type:varchar(100);not null;comment:请求方法"`
	RequestPath   string    `gorm:"type:varchar(100);not null;comment:请求路径"`
	IsDeleted     int       `gorm:"type:int;not null;default:0;comment:是否删除"`
}

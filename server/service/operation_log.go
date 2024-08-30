package service

import (
	"github.com/blkcor/gin-react-admin/core/db"
	"github.com/blkcor/gin-react-admin/models/model"
	"github.com/google/uuid"
)

// CreateOperationLog 创建操作日志
func CreateOperationLog(operator string, statusCode int, requestIp, requestMethod, requestPath string) error {
	return db.DB.Create(&model.OperationLog{
		ID:            uuid.New().ID(),
		Operator:      operator,
		StatusCode:    statusCode,
		RequestIP:     requestIp,
		RequestMethod: requestMethod,
		RequestPath:   requestPath,
	}).Error
}

// DeleteOperationLogRecord 删除操作日志
func DeleteOperationLogRecord(id uint32) error {
	return db.DB.Model(&model.OperationLog{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

// DeleteOperationLogByIds 根据id批量删除操作日志
func DeleteOperationLogByIds(ids []uint32) error {
	//将id_deleted设置为1
	return db.DB.Model(&model.OperationLog{}).Where("id in (?)", ids).Update("is_deleted", 1).Error
}

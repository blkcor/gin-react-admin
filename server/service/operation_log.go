package service

import (
	"github.com/blkcor/gin-react-admin/core/db"
	"github.com/blkcor/gin-react-admin/models/model"
	"github.com/blkcor/gin-react-admin/models/request"
	"github.com/blkcor/gin-react-admin/models/response"
	"github.com/blkcor/gin-react-admin/utils/pagination"
	"github.com/google/uuid"
	"strings"
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

// GetOperationLogList 获取操作日志列表
func GetOperationLogList(req request.GetOperationLogListRequest) (response.GetOperationLogListResponse, error) {
	resp := response.GetOperationLogListResponse{
		Paginator: pagination.Paginator{
			Size: req.Size,
			Page: req.Page,
		},
		Data: nil,
	}
	var operationLogs []model.OperationLog
	var total int64
	exec := db.DB.Model(&model.OperationLog{})
	if req.Path != "" {
		exec = exec.Where("request_path like ?", "%"+req.Path+"%")
	}
	if req.Method != "" {
		exec = exec.Where("request_method = ?", strings.ToUpper(req.Method))
	}
	if req.Ip != "" {
		exec = exec.Where("request_ip = ?", req.Ip)
	}
	err := exec.Count(&total).Error
	if err != nil {
		return resp, err
	}
	err = exec.Limit(req.Size).Offset((req.Page - 1) * req.Size).Find(&operationLogs).Error
	resp.Data = operationLogs
	resp.Total = total
	return resp, err
}

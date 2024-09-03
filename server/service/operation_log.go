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
	resp := response.GetOperationLogListResponse{}
	// 分页器
	p := pagination.Pagination{}
	p.Size = req.Size
	p.Page = req.Page

	var operationLogs []model.OperationLog
	exec := db.DB.Where("is_deleted = ?", 0)
	if req.Path != "" {
		exec = exec.Where("request_path like ?", "%"+req.Path+"%")
	}
	if req.Method != "" {
		exec = exec.Where("request_method = ?", strings.ToUpper(req.Method))
	}
	if req.Ip != "" {
		exec = exec.Where("request_ip = ?", req.Ip)
	}
	db.DB.Scopes(p.Paginate(operationLogs, db.DB)).Find(&operationLogs)
	resp.Data = operationLogs
	resp.Pagination = p
	return resp, nil
}

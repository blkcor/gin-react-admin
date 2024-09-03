package request

import "github.com/blkcor/gin-react-admin/utils/pagination"

// DeleteOperationLogByIdsRequest 根据id批量删除操作日志请求
type DeleteOperationLogByIdsRequest struct {
	Ids []uint32 `form:"ids" binding:"required"`
}

// GetOperationLogListRequest 获取操作日志列表请求
type GetOperationLogListRequest struct {
	pagination.Pagination
	Method string `json:"method"`
	Path   string `json:"path"`
	Ip     string `json:"ip"`
}

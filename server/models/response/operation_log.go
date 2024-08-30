package response

import (
	"github.com/blkcor/gin-react-admin/models/model"
	"github.com/blkcor/gin-react-admin/utils/pagination"
)

// GetOperationLogListResponse 获取操作日志列表响应
type GetOperationLogListResponse struct {
	Data []model.OperationLog `json:"data"`
	pagination.Paginator
}

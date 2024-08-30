package request

// DeleteOperationLogByIdsRequest 根据id批量删除操作日志请求
type DeleteOperationLogByIdsRequest struct {
	Ids []uint32 `form:"ids" binding:"required"`
}

package v1

import (
	"github.com/blkcor/gin-react-admin/core/logger"
	"github.com/blkcor/gin-react-admin/models/request"
	"github.com/blkcor/gin-react-admin/models/response"
	"github.com/blkcor/gin-react-admin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeleteOperationLogRecord 删除操作日志
func DeleteOperationLogRecord(context *gin.Context) {
	//获取参数
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(400, response.BaseResponse[any]{
			Success: false,
			Message: "参数错误",
			Data:    nil,
		})
		return
	}
	err = service.DeleteOperationLogRecord(uint32(idInt))
	if err != nil {
		logger.Error("删除操作日志记录失败: ", err)
		context.JSON(500, response.BaseResponse[any]{
			Success: false,
			Message: "删除操作日志记录失败",
			Data:    nil,
		})
		return
	}
	context.JSON(200, response.BaseResponse[any]{
		Success: true,
		Message: "删除操作日志记录成功",
		Data:    nil,
	})
}

// DeleteOperationLogByIds 批量删除操作日志
func DeleteOperationLogByIds(context *gin.Context) {
	//获取参数
	var deleteOperationLogRequest request.DeleteOperationLogByIdsRequest
	logger.Info("Content-Type: ", context.Request.Header.Get("Content-Type"))
	if err := context.ShouldBind(&deleteOperationLogRequest); err != nil {
		context.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Success: false,
			Message: "参数错误",
			Data:    nil,
		})
		return
	}
	err := service.DeleteOperationLogByIds(deleteOperationLogRequest.Ids)
	if err != nil {
		logger.Error("批量删除操作日志记录失败: ", err)
		context.JSON(500, response.BaseResponse[any]{
			Success: false,
			Message: "批量删除操作日志记录失败",
			Data:    nil,
		})
		return
	}
	context.JSON(200, response.BaseResponse[any]{
		Success: true,
		Message: "批量删除操作日志记录成功",
		Data:    nil,
	})
}

func GetOperationLogList(context *gin.Context) {
	//获取参数
	var req request.GetOperationLogListRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Success: false,
			Message: "参数错误",
			Data:    nil,
		})
		return
	}
	// 默认参数
	resp, err := service.GetOperationLogList(req)
	if err != nil {
		logger.Error("获取操作日志列表失败: ", err)
		context.JSON(500, response.BaseResponse[any]{
			Success: false,
			Message: "获取操作日志列表失败",
			Data:    nil,
		})
		return
	}
	context.JSON(200, response.BaseResponse[response.GetOperationLogListResponse]{
		Success: true,
		Message: "获取操作日志列表成功",
		Data:    resp,
	})
}

package v1

import (
	"github.com/blkcor/gin-react-admin/core/cache"
	"github.com/blkcor/gin-react-admin/models/model"
	"github.com/blkcor/gin-react-admin/models/response"
	"github.com/blkcor/gin-react-admin/service"
	"github.com/blkcor/gin-react-admin/utils/key"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetOnlineUser 获取在线用户列表
func GetOnlineUser(context *gin.Context) {
	members, err := cache.RDB.SMembers(context, key.OnlineUserSetKey).Result()
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Success: false,
			Message: "获取在线用户失败",
			Data:    nil,
		})
		return
	}
	userIdList := make([]uint32, 0)
	for _, member := range members {
		userId := member[len(key.OnlineUserPrefix):]
		userIdU32, _ := strconv.Atoi(userId)
		userIdList = append(userIdList, uint32(userIdU32))
	}

	userList, err := service.GetUserByIds(userIdList)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Success: false,
			Message: "获取在线用户失败",
			Data:    nil,
		})
		return
	}
	context.JSON(http.StatusOK, response.BaseResponse[[]model.User]{
		Success: true,
		Message: "获取在线用户成功",
		Data:    userList,
	})
}

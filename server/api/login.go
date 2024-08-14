package api

import (
	"github.com/blkcor/gin-react-admin/config/section"
	"github.com/blkcor/gin-react-admin/models/request"
	"github.com/blkcor/gin-react-admin/models/response"
	"github.com/blkcor/gin-react-admin/service"
	"github.com/blkcor/gin-react-admin/utils/captcha"
	"github.com/blkcor/gin-react-admin/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 登录
func Login(context *gin.Context) {
	loginRequest := request.LoginRequest{}
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	//校验用户的合法性
	user, err := service.GetUserByUsername(loginRequest.Username)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "用户不存在",
		})
		return
	}
	//校验密码的合法性
	valid := user.CheckPassword(loginRequest.Password)
	if !valid {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "密码错误",
		})
		return
	}
	//校验验证码 对比redis和用户输入的验证码是否一致
	ok := captcha.VerifyCaptcha(context, loginRequest.Captcha)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "验证码错误",
		})
		return
	}
	//颁发token
	//获取用户的角色信息
	role, err := service.GetUserRoleByUserId(user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取用户角色信息失败",
		})
		return
	}
	token, err := jwt.GenToken(user.ID, user.Username, user.Email, role.RoleCode, section.AppConfig.AccessKey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "生成token失败",
		})
		return
	}
	context.JSON(http.StatusOK, response.LoginResponse{
		Token:   token,
		Message: "登录成功",
		User: response.UserInfo{
			UserId:   user.ID,
			Username: user.Username,
			Email:    user.Email,
			Avatar:   user.Avatar,
		},
	})
}

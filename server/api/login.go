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

// Login godoc
//
// @Summary      登录
// @Description  管理登录接口，用户可以通过该接口进行登录。登录时需要提供用户名、密码和验证码。接口会校验用户的合法性、密码和验证码，如果校验成功，则生成 JWT token 并返回。如果校验失败，则返回相应的错误信息。
// @Tags         Login
// @Accept       json
// @Produce      json
// @Param        body  body     request.LoginRequest  true  "登录请求参数"
// @Success      200  {object}  response.LoginResponse  "登录成功，返回用户信息和 JWT token"
// @Failure      400  {object}  map[string]interface{}  "参数错误，返回详细错误信息"
// @Failure      401  {object}  map[string]interface{}  "用户认证失败，包括用户名不存在、密码错误或验证码错误"
// @Failure      500  {object}  map[string]interface{}  "服务器内部错误，返回详细错误信息"
// @Security     ApiKeyAuth
// @Router       /login [post]
func Login(context *gin.Context) {
	loginRequest := request.LoginRequest{}
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	// 校验用户的合法性
	user, err := service.GetUserByUsername(loginRequest.Username)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "用户不存在",
		})
		return
	}
	// 校验密码的合法性
	valid := user.CheckPassword(loginRequest.Password)
	if !valid {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "密码错误",
		})
		return
	}
	// 校验验证码，对比 Redis 和用户输入的验证码是否一致
	ok := captcha.VerifyCaptcha(context, loginRequest.Captcha)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "验证码错误",
		})
		return
	}
	// 颁发 token
	// 获取用户的角色信息
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
			"message": "生成 token 失败",
		})
		return
	}
	context.JSON(http.StatusOK, response.LoginResponse{
		Token:   token,
		Message: "登录成功",
		UserInfo: response.UserInfo{
			UserId:   user.ID,
			Username: user.Username,
			Email:    user.Email,
			Avatar:   user.Avatar,
		},
	})
}

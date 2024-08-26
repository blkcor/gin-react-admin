package response

// LoginResponse 登录响应
type LoginResponse struct {
	BaseResponse[UserInfo]
	Token string `json:"token"`
}

// LogoutResponse 登录响应
type LogoutResponse struct {
	BaseResponse[any]
}

// UserInfo 用户信息
type UserInfo struct {
	UserId   uint32 `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	UserRole string `json:"userRole"`
	RoleCode string `json:"roleCode"`
}

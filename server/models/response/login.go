package response

type LoginResponse struct {
	Token    string   `json:"token"`
	Message  string   `json:"message"`
	UserInfo UserInfo `json:"userInfo"`
}

type UserInfo struct {
	UserId   uint32 `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

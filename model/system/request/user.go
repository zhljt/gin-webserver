package request

// 登陆请求参数
type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// 注册请求参数
type RegisterRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// 修改密码请求参数
type ChangePasswordRequest struct {
	UserId      uint   `json:"userId"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newpassword"`
}

// 重置密码请求参数
type ResetPasswordRequest struct {
	UserId []uint `json:"userId"`
}

// 获取用户信息请求参数
type GetUserInfoRequest struct {
	UserId uint `json:"userId"`
}

// 更新用户信息请求参数
type UpdateUserInfoRequest struct {
	UserId   uint   `json:"userId"`
	Account  string `json:"account"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// 修改用户状态请求参数,0-锁定,1-正常,2-删除
type ChangeUserStatusRequest struct {
	UserId uint  `json:"userId"`
	Status uint8 `json:"status"`
}

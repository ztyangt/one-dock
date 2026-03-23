package user

type getUserRequest struct {
	ID int64 `query:"id" validate:"required" alias:"用户id"`
}

// loginRequest 登录请求体
type loginRequest struct {
	Account  string `json:"account" validate:"required" alias:"账号"`
	Password string `json:"password" validate:"required" alias:"密码"`
}

// userInfoResponse 用户信息响应体
type userInfoResponse struct {
	ID          string `json:"id"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Account     string `json:"account"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Gender      int64  `json:"gender"`
}

// loginResponse 登录响应体
type loginResp struct {
	Token string           `json:"token"`
	User  userInfoResponse `json:"user"`
}

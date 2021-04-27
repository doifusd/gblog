package request

//AuthRequest jwt 请求
type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CheckUserRequest struct {
	Name string `form:"name" binding:"required"`
}

type LoginRequest struct {
	Mobile   string `json:"mobile" binding:"required,min=11,max=11"`
	Password string `json:"password" binding:"required"`
}

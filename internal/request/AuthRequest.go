package request

//AuthRequest jwt 请求
type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

type RegisterRequest struct {
	Name   string `form:"name" binding:"required"`
	Mobile string `form:"mobile" binding:"required"`
	Passwd string `form:"password" binding:"required"`
}

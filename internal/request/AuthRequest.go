package request

//AuthRequest jwt 请求
type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"reuqired"`
	AppSecret string `form:"app_secret" binding:"required"`
}

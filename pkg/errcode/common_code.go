package errcode

var (
	Success                   = NewERror(0, "成功")
	ServerError               = NewERror(100000000, "服务器内部错误")
	IntvalidParams            = NewERror(100000001, "入参错误")
	NotFound                  = NewERror(100000002, "找不到")
	UnauthorizedAuthNotExist  = NewERror(100000003, "鉴权失败，找不到对应的appkey和appsecret")
	UnauthorizedTokenError    = NewERror(100000004, "鉴权失败，token错误")
	UnauthorizedTokenTimeout  = NewERror(100000005, "鉴权失败，token超时")
	UnauthorizedTokenGenerate = NewERror(100000006, "鉴权失败，token生成失败")
	TooManyRequest            = NewERror(100000007, "请求过多")
)

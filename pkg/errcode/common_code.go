package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(100000000, "服务器内部错误")
	IntvalidParams            = NewError(100000001, "入参错误")
	NotFound                  = NewError(100000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(100000003, "鉴权失败，找不到对应的appkey和appsecret")
	UnauthorizedTokenError    = NewError(100000004, "鉴权失败，token错误")
	UnauthorizedTokenTimeout  = NewError(100000005, "鉴权失败，token超时")
	UnauthorizedTokenGenerate = NewError(100000006, "鉴权失败，token生成失败")
	TooManyRequest            = NewError(100000007, "请求过多")
	LimiterRequest            = NewError(100000008, "限流失败")
)

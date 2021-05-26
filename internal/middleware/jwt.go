package middleware

import (
	"blog/global"
	"blog/pkg/app"
	"blog/pkg/errcode"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.IntvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			resp := app.NewResponse(c)
			resp.ToErrorResponse(ecode)
			//尽快结束当前请求
			c.Abort()
			return
		}
		uid, err := global.Redis.Get(c, token+":uid").Result()
		if err != nil {
			panic(err)
		}
		c.Set("uid", uid)
		c.Next()
	}
}

package middleware

import (
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/limiter"

	"github.com/gin-gonic/gin"
)

//RaleLimiter 中间件限流
func RaleLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				resp := app.NewResponse(c)
				resp.ToErrorResponse(errcode.TooManyRequest)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

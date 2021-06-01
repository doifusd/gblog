package middleware

import (
	"blog/global"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/limiter"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

var tb_cache_time string = "10"
var token_bucket_size int64 = 10
var tb_time_key string = "time_key"
var tb_value string = "1"
var tb_counter_key string = "counter_key"

func RedisLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		e, err := global.Redis.Get(c, tb_time_key).Result()
		if err != nil {
			log.Println(err)
		}
		if e == "" {
			luaId := redis.NewScript(`
			local time_key = KEYS[1]
			local time_key_timeout = KEYS[2]
			local cnt_key = KEYS[3]
			local cnt_value = KEYS[4]
			local r1 = redis.call('setex',time_key,time_key_timeout,cnt_value)
			if r1 == "ok"
			then
			local r2 = redis.call('setex',cnt_key,time_key_timeout,cnt_value)
			return r2
			else
			return -1
			end
			`)
			keys := []string{tb_time_key, tb_cache_time, tb_counter_key, tb_value}
			n, err := luaId.Run(c, global.Redis, keys, 4).Result()
			if err != nil {
				log.Println(err)
			}
			if n == -1 {
				resp := app.NewResponse(c)
				resp.ToErrorResponse(errcode.LimiterRequest)
				c.Abort()
				return
			}
		} else {
			luaId := redis.NewScript(`
			local cnt_key = KEYS[1]
			local time_key = KEYS[2]
			local current_cnt = redis.call('incr',cnt_key)
			local t = redis.call('ttl',time_key)
			redis.call('expire',cnt_key,t)
			return current_cnt
			`)
			keys := []string{tb_counter_key, tb_time_key}
			n, err := luaId.Run(c, global.Redis, keys, 2).Result()
			if err != nil {
				log.Println(err)
				resp := app.NewResponse(c)
				resp.ToErrorResponse(errcode.LimiterRequest)
				c.Abort()
			}
			if n == -1 {
				resp := app.NewResponse(c)
				resp.ToErrorResponse(errcode.LimiterRequest)
				c.Abort()
				return
			}

			if n.(int64) > token_bucket_size {
				//调用频率过快
				resp := app.NewResponse(c)
				resp.ToErrorResponse(errcode.TooManyRequest)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

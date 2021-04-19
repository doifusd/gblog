package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

//LimiterIface 限流接口
type LimiterIface interface {
	Key(c *gin.Context) string                         //追踪限流器的键值对的名称
	GetBucket(key string) (*ratelimit.Bucket, bool)    //获取令牌桶
	AddBucket(rules ...LimiterBucketRule) LimiterIface //新增多个令牌桶
}

//Limiter 限流结构体
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

//LimiterBucketRule 限流规则
type LimiterBucketRule struct {
	Key          string        //自定义键值对名称
	FillInterval time.Duration //间隔多久时间防n个令牌
	Capacity     int64         //令牌桶的容量
	Quantum      int64         //每次到达间隔事件后所放的具体令牌数量
}

package routers

import (
	"blog/global"
	"blog/internal/middleware"
	"blog/internal/routers/api"
	v1 "blog/internal/routers/api/v1"
	"blog/pkg/limiter"
	"net/http"
	"time"

	_ "blog/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//限流器
var methodLimiters = limiter.NewMethodLimiter().AddBucket(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

//NewRoter 初始化路由器
func NewRoter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RaleLimiter(methodLimiters))
	//请求超时设置
	//r.Use(middleware.ContextTimeout(60*time.Second))
	r.Use(middleware.ContextTimeout(global.AppSetting.RequestTimeout * time.Second))
	r.Use(middleware.Translations())

	r.Use(middleware.Tracer())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	upload := NewUpload()
	r.POST("upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UpLoadSavePath))

	article := v1.NewArticle()
	tag := v1.NewTag()

	r.GET("/auth", api.GetAuth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.List)

	}
	return r
}

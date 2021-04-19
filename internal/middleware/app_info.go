package middleware

import (
	"blog/global"

	"github.com/gin-gonic/gin"
)

//AppInfo 服务信息
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", global.AppSetting.AppName)
		c.Set("app_version", global.AppSetting.AppVersion)
		// c.Set("app_name", "blog")
		// c.Set("app_version", "1.0.0")
		c.Next()
	}
}

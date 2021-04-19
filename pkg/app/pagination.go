package app

import (
	"blog/global"
	"blog/pkg/convert"

	"github.com/gin-gonic/gin"
)

//GetPage 分页获取页数
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

//GetPageSize 每页大小
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

//GetPageOffset offset大小
func GetPageOffset(page, pageSize int) int {
	res := 0
	if page > 0 {
		res = (page - 1) * pageSize
	}
	return res
}

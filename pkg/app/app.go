package app

import (
	"blog/pkg/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Response 响应结构体
type Response struct {
	Ctx *gin.Context
}

//Pager 分页结构体
type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

//NewResponse 实例化响应
func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

//ToResponse 响应方法
func (r *Response) ToResponse(data interface{}) {
	// resp := gin.H{"code": err.Code(), "msg": err.Msg()}
	// details := err.Details()
	if data == nil {
		data = gin.H{}
	}
	// if len(details) > 0 {
	// resp["details"] = details
	// }
	r.Ctx.JSON(http.StatusOK, data)
}

//ToResponseList 响应列表
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

//ToErrorResponse 错误响应
func (r *Response) ToErrorResponse(err *errcode.Error) {
	resp := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		resp["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), resp)
}

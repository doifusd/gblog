package v1

import (
	"blog/global"
	"blog/internal/request"
	"blog/internal/service"
	"blog/pkg/app"
	"blog/pkg/convert"
	"blog/pkg/errcode"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query string false "状态" Enums(0,1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := request.TagListRequest{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.IntvalidParams.WithDetails(errs.Errors()...)
		resp.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&request.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag errs: %v", errs)
		resp.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTaglist errs: %v", err)
		resp.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	resp.ToResponseList(tags, int(totalRows))
	return
}

// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := request.CrateTagResquest{}
	resp := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", err)
		errRsp := errcode.IntvalidParams.WithDetails(err.Errors()...)
		resp.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	//检测数据是否存在
	checkTag, errs := svc.GetTag(&param)
	if errs != nil {
		global.Logger.Errorf("svc.CreateTag errs: %v", errs)
		resp.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	fmt.Println(checkTag)
	if checkTag > 0 {
		resp.ToErrorResponse(errcode.ErrorTagExist)
		return
	}

	errss := svc.CreateTag(&param)
	if errss != nil {
		global.Logger.Errorf("svc.CreateTag errs: %v", errss)
		resp.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	resp.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce json
// @Param id path int true "标签ID"
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := request.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	resp := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &param)
	fmt.Println(valid)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", err)
		errRsp := errcode.IntvalidParams.WithDetails(err.Errors()...)
		resp.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	errs := svc.UpdateTag(&param)
	if errs != nil {
		global.Logger.Errorf("svc.UpdateTag errs: %v", errs)
		resp.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	resp.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := request.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	resp := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", err)
		errRsp := errcode.IntvalidParams.WithDetails(err.Errors()...)
		resp.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	errs := svc.DeleteTag(&param)
	if errs != nil {
		global.Logger.Errorf("svc.UpdateTag errs: %v", errs)
		resp.ToErrorResponse(errcode.ErrorDeleateTagFail)
		return
	}
	resp.ToResponse(gin.H{})
	return
}

func (t Tag) Get(c *gin.Context) {}

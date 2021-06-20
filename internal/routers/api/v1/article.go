package v1

import (
	"blog/global"
	"blog/internal/request"
	"blog/internal/service"
	"blog/pkg/app"
	"blog/pkg/convert"
	"blog/pkg/errcode"
	"blog/pkg/qrcode"
	"time"

	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
)

const (
	QRCODE_URL = "https://github.com/EDDYCJY/blog#gin%E7%B3%BB%E5%88%97%E7%9B%AE%E5%BD%95"
)

type Article struct{}

type AticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取文章详情
// @Produce json
// @Param id query string false "文章id"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/aritlce/:id [get]
func (a Article) Info(c *gin.Context) {
	start_time := time.Now().UnixNano()
	resp := app.NewResponse(c)
	param := request.ArticleInfoRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.IntvalidParams.WithDetails(errs.Errors()...)
		resp.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	article, err := svc.GetArticleInfo(&param)
	if err != nil {
		global.Logger.Errorf("svc.ArticleInfo errs: %v", err)
		resp.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}
	stop_time := time.Now().UnixNano()
	e_time := (stop_time - start_time) / 1e6
	data := gin.H{"code": errcode.SuccessGetArticle.Code(), "msg": errcode.SuccessGetArticle.Msg(), "info": article, "e_time": e_time}
	resp.ToResponse(data)
	return
}

// @Summary 获取文章列表
// @Produce json
// @Param created_by query string false "创建者id"
// @Param state query string false "状态" Enums(0,1) default(1)
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/aritlces [get]
func (a Article) List(c *gin.Context) {
	start_time := time.Now().UnixNano()
	resp := app.NewResponse(c)
	param := request.ArticleListRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.IntvalidParams.WithDetails(errs.Errors()...)
		resp.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountArtilce(&request.CountArticleRequest{CreatedBy: param.CreatedBy, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountArticle errs: %v", errs)
		resp.ToErrorResponse(errcode.ErrorCountArticleFail)
		return
	}
	tags, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetArticlelist errs: %v", err)
		resp.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}
	stop_time := time.Now().UnixNano()
	e_time := (stop_time - start_time) / 1e6
	resp.ToResponseList(tags, int(totalRows), e_time)
	return
}

// @Summary 新增文章
// @Produce json
// @Param title body string true "文章标题" minlength(3) maxlength(100)
// @Param desc body string true "描述" minlength(3) maxlength(100)
// @Param content body string true "内容" minlength(3) maxlength(1000)
// @Param created_by body int true "创建者"
// @Param cover body string false "封面图"
// @Param tags body []int  "使用标签"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article [post]
func (a Article) Create(c *gin.Context) {
	param := request.CreateArticleResquest{}
	resp := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", err)
		errRsp := errcode.IntvalidParams.WithDetails(err.Errors()...)
		resp.ToErrorResponse(errRsp)
		return
	}
	param.CreatedBy = 0
	uid, exit := c.Get("uid")
	if exit == true {
		param.CreatedBy = convert.StrTo(uid.(string)).MustUInt32()
	}
	svc := service.New(c.Request.Context())
	errss := svc.CreateArticle(&param)
	if errss != nil {
		global.Logger.Errorf("svc.CreateArticle errs: %v", errss)
		resp.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	resp.ToSuccessResponse(errcode.SuccessCreateArticle)
	return
}

// @Summary 编辑文章
// @Produce json
// @Param title body string true "文章标题" minlength(3) maxlength(100)
// @Param desc body string true "描述" minlength(3) maxlength(100)
// @Param content body string true "内容" minlength(3) maxlength(1000)
// @Param modified_by body int true "修改者"
// @Param state body true int "状态"
// @Param cover body string false "封面图"
// @Param tags body []int  "使用标签"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := request.UpdateArticleRequest{}
	resp := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", err)
		errRsp := errcode.IntvalidParams.WithDetails(err.Errors()...)
		resp.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())

	param.ModifiedBy = 0
	uid, exit := c.Get("uid")
	if exit == true {
		param.ModifiedBy = convert.StrTo(uid.(string)).MustUInt32()
	}
	errss := svc.UpdateArticle(&param)
	if errss != nil {
		global.Logger.Errorf("svc.updateArticle errs: %v", errss)
		resp.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	resp.ToSuccessResponse(errcode.SuccessUpdateArticle)
	return
}

// @Summary 删除文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {array} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := request.DeleteArticleRequest{
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

	param.ModifiedBy = 0
	uid, exit := c.Get("uid")
	if exit == true {
		param.ModifiedBy = convert.StrTo(uid.(string)).MustUInt32()
	}
	errs := svc.DeleteArticle(&param)
	if errs != nil {
		global.Logger.Errorf("svc.delArticle errs: %v", errs)
		resp.ToErrorResponse(errcode.ErrorDeleateArticleFail)
		return
	}
	resp.ToSuccessResponse(errcode.SuccessDeleteArticle)
	return
}

func (a Article) GenerateArticlePoster(c *gin.Context) {
	resp := app.NewResponse(c)
	qrc := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto)
	path := qrcode.GetQrCodeFullPath()
	_, _, err := qrc.Encode(path)
	if err != nil {
		resp.ToSuccessResponse(errcode.SuccessGetArticle)
		return
	}

	resp.ToSuccessResponse(errcode.SuccessGetArticle)
	return
}

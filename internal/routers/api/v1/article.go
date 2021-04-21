package v1

import (
	"blog/pkg/app"
	"blog/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Article struct{}

type AticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

/*
func (a Article) List(c *gin.Context) {
	param := request.ArticleListRequest{}
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
	totalRows, err := svc.CountArticle(&request.CountArticleRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag errs: %v", errs)
		resp.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTaglist errs: %v", err)
		resp.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	resp.ToResponseList(tags, int(totalRows))
	return
}

func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
*/

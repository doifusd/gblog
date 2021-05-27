package service

import (
	"blog/internal/model"
	"blog/internal/request"
	"blog/pkg/app"
)

func (svc *Service) CountArtilce(param *request.CountArticleRequest) (int64, error) {
	return svc.dao.CountArticle(param.CreatedBy, param.State)
}

func (svc *Service) GetArticleInfo(param *request.ArticleInfoRequest) (*model.Article, error) {
	return svc.dao.GetArticleInfo(param.ArticleID)
}

func (svc *Service) GetArticleList(param *request.ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.CreatedBy, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateArticle(param *request.CreateArticleResquest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.Cover, param.CreatedBy, param.Tags)
}

func (svc *Service) UpdateArticle(param *request.UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.Cover, param.State, param.ModifiedBy, param.Tags)
}

func (svc *Service) DeleteArticle(param *request.DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID, param.ModifiedBy)
}

/*
func (svc *Service) GetArticle(param *request.CrateArticleResquest) (int64, error) {
	return svc.dao.GetArticle(param.Title, param.CreatedBy)
}
*/

package service

import (
	"blog/internal/model"
	"blog/internal/request"
	"blog/pkg/app"
)

func (svc *Service) CountTag(param *request.CountTagRequest) (int64, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagInfo(param *request.TagInfoRequest) (*model.Tag, error) {
	return svc.dao.GetTagInfo(param.TagId)
}

func (svc *Service) GetTagList(param *request.TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) GetTag(param *request.CreateTagResquest) (int64, error) {
	return svc.dao.GetTag(param.Name, param.CreatedBy)
}

func (svc *Service) CreateTag(param *request.CreateTagResquest) error {
	return svc.dao.CreateTag(param.Name, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *request.UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *request.DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID, param.ModifiedBy)
}

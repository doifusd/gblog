package service

import (
	"blog/internal/model"
	"blog/internal/request"
)

func (svc *Service) CheckUser(param *request.LoginRequest) (*model.User, error) {
	return svc.dao.GetUser(param.Mobile)
}

func (svc *Service) CreateUserCheck(name string) error {
	return svc.dao.GetUserName(name)
}

func (svc *Service) CreateUser(param *request.RegisterRequest) error {
	return svc.dao.CreateUser(param.Name, param.Mobile, param.Password)
}

func (svc *Service) UserLastId() (uint32, error) {
	return svc.dao.LastId()
}

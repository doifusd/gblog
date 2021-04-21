package service

import (
	"blog/internal/model"
	"blog/internal/request"
)

func (svc *Service) CheckUser(param *request.LoginRequest) (*model.User, error) {
	return svc.dao.GetUser(param.Mobile)
}

func (svc *Service) CreateUser(param *request.RegisterRequest) error {
	return svc.dao.CreateUser(param.Name, param.Mobile, param.Passwd)
}

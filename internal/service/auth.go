package service

import (
	"blog/internal/request"
	"errors"
)

//CheckAuth 认证检测
func (svc *Service) CheckAuth(param *request.AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist.")
}

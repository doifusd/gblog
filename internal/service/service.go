package service

import (
	"blog/global"
	"blog/internal/dao"
	"blog/internal/model"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(model.WithContext(svc.ctx, global.DBEngine))
	// svc.dao = dao.New(global.DBEngine)
	return svc
}

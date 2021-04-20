package service

import (
	"blog/global"
	"blog/internal/dao"
	"blog/pkg/tracer"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(tracer.WithContext(svc.ctx, global.DBEngine))
	// svc.dao = dao.New(global.DBEngine)
	return svc
}

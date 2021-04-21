package api

import (
	"blog/global"
	"blog/internal/request"
	"blog/internal/service"
	"blog/pkg/app"
	"blog/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	param := request.LoginRequest{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		errResp := errcode.IntvalidParams.WithDetails(errs.Errors()...)
		resp.ToErrorResponse(errResp)
		return
	}
	svc := service.New(c.Request.Context())
	user, err := svc.CheckUser(&param)
	if err != nil {
		// global.Logger.Errorf("svc.CheckAuth err: %v", err)
		// resp.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	// util.AESCBCEncrypt(param.Passwd,""+"id")
	if user.Passwd != param.Passwd {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		resp.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.Mobile, param.Passwd)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		resp.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	resp.ToResponse(gin.H{
		"token": token,
	})
}

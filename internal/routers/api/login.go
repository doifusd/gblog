package api

import (
	"blog/global"
	"blog/internal/request"
	"blog/internal/service"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/util"
	"fmt"

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
		global.Logger.Errorf("svc.CheckUsername err: %v", err)
		resp.ToErrorResponse(errcode.ErrorUserNotExist)
		return
	}
	param.Password = util.UserPassword(param.Password, fmt.Sprintf("%d", user.ID))
	if user.Passwd != param.Password {
		global.Logger.Errorf("user.password err: %v", err)
		resp.ToErrorResponse(errcode.ErrorUserPasswordFail)
		return
	}
	token, err := app.GenerateToken(param.Mobile, param.Password)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		resp.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	//存储用户信息
	data := make(map[string]string, 1)
	data["token"] = token
	resp.ToResponse(gin.H{
		"code": "100030008",
		"msg":  "登录成功",
		"data": data,
		// "token": token,
	})
}

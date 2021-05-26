package api

import (
	"blog/global"
	"blog/internal/request"
	"blog/internal/service"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var cacheTime time.Duration

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
	uidStr := fmt.Sprintf("%d", user.ID)

	param.Password = util.UserPassword(param.Password, uidStr)
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
	cacheTime = 86400 * 15 * time.Second
	err = global.Redis.Set(c, token+":uid", user.ID, cacheTime).Err()
	if err != nil {
		panic(err)
	}

	data := make(map[string]string, 1)
	data["token"] = token

	resp.ToResponse(gin.H{
		"code": "0",
		"msg":  "登录成功",
		"data": data,
	})
}

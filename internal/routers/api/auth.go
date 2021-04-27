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

/*
* GetAuth 生成token
func GetAuth(c *gin.Context) {
	param := request.AuthRequest{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		errResp := errcode.IntvalidParams.WithDetails(errs.Errors()...)
		resp.ToErrorResponse(errResp)
		return
	}
	svc := service.New(c.Request.Context())
	chek,err := svc.CheckUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		resp.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		resp.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	resp.ToResponse(gin.H{
		"token": token,
	})
}*/

func SignUp(c *gin.Context) {
	param := request.RegisterRequest{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		errResp := errcode.IntvalidParams.WithDetails(errs.Errors()...)
		resp.ToErrorResponse(errResp)
		return
	}
	svc := service.New(c.Request.Context())
	//检查mobile是否存在
	err := svc.CreateUserCheck(param.Name)
	if err == nil {
		resp.ToErrorResponse(errcode.ErrorUserExist)
		return
	}
	lastId, err := svc.UserLastId()
	lastIds := "1"
	if err != nil {
		lastIds = fmt.Sprintf("%d", lastId+1)
	}
	// var build strings.Builder
	// build.WriteString(global.AppSetting.AppName)
	// build.WriteString(lastIds)
	// key := build.String()
	// tmps := util.AESCBCEncrypt([]byte(param.Password), []byte(key))
	// encodedStr := hex.EncodeToString(tmps)
	// param.Password = string(encodedStr)
	param.Password = util.UserPassword(param.Password, lastIds)
	err = svc.CreateUser(&param)
	if err != nil {
		resp.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}
	// resp.ToResponse(gin.H{"code": "100030001", "msg": "注册成功"})
	resp.ToErrorResponse(errcode.SuccessCreateUser)
}

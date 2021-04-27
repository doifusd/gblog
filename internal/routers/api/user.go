package api

/*func checkUserName(c *gin.Context) {
	param := request.CheckUserRequest{}
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
	err := svc.GetUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		resp.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	resp.ToErrorResponse(errcode.SuccessCreateUser)
}*/

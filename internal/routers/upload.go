package routers

import (
	"blog/global"
	"blog/internal/service"
	"blog/pkg/app"
	"blog/pkg/convert"
	"blog/pkg/errcode"
	"blog/pkg/upload"
	"blog/pkg/util"
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	resp := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errsp := errcode.IntvalidParams.WithDetails(err.Error())
		resp.ToErrorResponse(errsp)
		return
	}
	if fileHeader == nil || fileType <= 0 {
		resp.ToErrorResponse(errcode.IntvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UPloadFile err: %v", err)
		errsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		resp.ToErrorResponse(errsp)
		return
	}
	resp.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}

func (u Upload) UploadFileMuli(c *gin.Context) {
	resp := app.NewResponse(c)
	//最大上传大小
	err := c.Request.ParseMultipartForm(32 << 20)
	//超过上传大小
	if err != nil {
		errsp := errcode.IntvalidParams.WithDetails(err.Error())
		resp.ToErrorResponse(errsp)
		return
	}
	fhs := c.Request.MultipartForm.File["file"]
	allowExt := global.AppSetting.UploadImageAllowExts
	for _, fheader := range fhs {
		// saveUploadImage(fheader)
		fmt.Println(fheader.Filename)
		//获取文件类型
		fileExt := path.Ext(fheader.Filename)
		fmt.Println(fileExt)
		if util.InArray(fileExt, allowExt) == false {
			resp.ToErrorResponse(errcode.ErrorUploadExtFail)
			return
		}
		//
	}

	// svc := service.New(c.Request.Context())
	//fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	//if err != nil {
	//	global.Logger.Errorf("svc.UPloadFile err: %v", err)
	//	errsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
	//	resp.ToErrorResponse(errsp)
	//	return
	//}
	//resp.ToResponse(gin.H{
	//	"file_access_url": fileInfo.AccessUrl,
	//})

	/*
		file, fileHeader, err := c.Request.FormFile("file")
		fileType := convert.StrTo(c.PostForm("type")).MustInt()
		if err != nil {
			errsp := errcode.IntvalidParams.WithDetails(err.Error())
			resp.ToErrorResponse(errsp)
			return
		}
		if fileHeader == nil || fileType <= 0 {
			resp.ToErrorResponse(errcode.IntvalidParams)
			return
		}

		svc := service.New(c.Request.Context())
		fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
		if err != nil {
			global.Logger.Errorf("svc.UPloadFile err: %v", err)
			errsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
			resp.ToErrorResponse(errsp)
			return
		}
		resp.ToResponse(gin.H{
			"file_access_url": fileInfo.AccessUrl,
		})
	*/
}

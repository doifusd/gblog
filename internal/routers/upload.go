package routers

import (
	"blog/global"
	"blog/internal/service"
	"blog/pkg/app"
	"blog/pkg/convert"
	"blog/pkg/errcode"
	"blog/pkg/upload"

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

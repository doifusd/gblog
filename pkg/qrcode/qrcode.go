package qrcode

import (
	"blog/global"
	"blog/pkg/file"
	"blog/pkg/util"
	"image/jpeg"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

//图片格式
const (
	EXT_JPG = ".jpg"
	//图片质量
	Qualitys = 85
)

//生成二维码
func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}
func GetQrCodePath() string {
	return global.AppSetting.QrCodeSavePath
}
func GetQrCodeFullPath() string {
	return global.AppSetting.QrCodeSavePath
}
func GetQrCodeFullUrl(name string) string {
	return global.AppSetting.PrefixUrl + "/" + GetQrCodePath() + name
}
func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}
func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}
func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	if file.CheckNotExist(src) == true {
		return false
	}
	return true
}
func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	if file.CheckNotExist(src) == true {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}
		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}
		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()
		err = jpeg.Encode(f, code, &jpeg.Options{Quality: Qualitys})
		if err != nil {
			return "", "", err
		}
	}
	return name, path, nil
}

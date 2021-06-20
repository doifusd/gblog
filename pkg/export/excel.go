package export

import "blog/global"

func GetExcelFullUrl(name string) string {
	return global.AppSetting.PrefixUrl + name
}
func GetExcelFullPath() string {
	return global.AppSetting.ExportSavePath
}

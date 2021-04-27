package util

import (
	"blog/global"
	"encoding/hex"
	"strings"
)

func UserPassword(Password, uid string) string {
	var build strings.Builder
	build.WriteString(global.AppSetting.AppName)
	build.WriteString(uid)
	key := build.String()
	tmps := AESCBCEncrypt([]byte(Password), []byte(key))
	encodedStr := hex.EncodeToString(tmps)
	return string(encodedStr)
}

package global

import (
	"blog/pkg/logger"
	"blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	CacheSetting    *setting.CacheSettings
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettings
	EmailSetting    *setting.EmailSettings
	TracerSetting   *setting.TracerSettings
)

package global

import (
	"mall_go/pkg/logger"
	"mall_go/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	EmailSetting    *setting.EmailSettings
	JWTSetting      *setting.JWTSettings
	DatabaseSetting *setting.DatabaseSettings
	RedisSetting    *setting.RedisSettings
	LogSetting      *setting.LogSettings
	Logger          *logger.Logger
)
var (
	ApiAuthConfig = map[string]map[string]string{
		// 调用方
		"DEMO": {
			"md5": "IgkibX71IEf382PT",
			"aes": "IgkibX71IEf382PT",
			"rsa": "rsa/public.pem",
		},
	}
	AppSignExpiry = "120"
)

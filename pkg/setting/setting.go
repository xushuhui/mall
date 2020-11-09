package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type ServerSettings struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Domain       string
}

type AppSettings struct {
	AppName               string
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	UploadSavePath        string
	UploadServerUrl       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
}

type EmailSettings struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

type JWTSettings struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type DatabaseSettings struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
type RedisSettings struct {
	Host     string
	Password string
}
type LogSettings struct {
	Level        string
	Formatter    string
	ReportCaller bool
	SavePath     string
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	e := s.vp.UnmarshalKey(k, v)
	if e != nil {
		return e
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		e := s.ReadSection(k, v)
		if e != nil {
			return e
		}
	}

	return nil
}

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configDir, configName, configType string) (*Setting, error) {
	vp := viper.New()
	vp.AddConfigPath(configDir)
	vp.SetConfigName(configName)

	vp.SetConfigType(configType)
	e := vp.ReadInConfig()
	if e != nil {
		return nil, e
	}

	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}

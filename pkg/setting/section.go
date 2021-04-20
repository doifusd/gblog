package setting

import "time"

//ServerSettings 服务配置
type ServerSettings struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// AppSettings 项目配置
type AppSettings struct {
	AppName              string
	AppVersion           string
	DefaultPageSize      int
	MaxPageSize          int
	RequestTimeout       time.Duration
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UpLoadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}

//DatabaseSettings 数据库配置
type DatabaseSettings struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	Port         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

//JWTSettings jwt 配置
type JWTSettings struct {
	Secret string
	Issuer string
	Expire time.Duration
}

//EmailSettings 配置
type EmailSettings struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

type TracerSettings struct {
	ServiceName   string
	AgentHostPort string
}

var sections = make(map[string]interface{})

//ReadSection 读取配置
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

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

type CacheSettings struct {
	Addr         string
	DB           int           //默认数据库
	Password     string        //密码
	PoolSize     int           //连接池最大socket连接数,默认为4倍CPU数,4*runtime.NumCPU
	MinIdleConns int           //在启动阶段创建指定数量的Idle连接,并长期维持idle状态的连接数
	DialTimeout  time.Duration //连接建立超时时间,默认5秒
	ReadTimeout  time.Duration //读超时,默认3秒,-1表示取消读超时
	WriteTimeout time.Duration //写超时,默认等于读超时
	PoolTimeout  time.Duration //客户端等待可用连接的最>大等待时长,默认为读超时+1秒
	//闲置连接检查包括IdleTimeout，MaxConnAge
	//闲置连接检查的周期,默认为1分钟,-1表示不做周期性检查，
	//只在客户端获取连接时对闲置连接进行处理。
	IdleCheckFrequency time.Duration
	IdleTimeout        time.Duration //闲置超时,默认5分钟,-1表示取消闲置超时检查
	MaxConnAge         time.Duration //连接存活时长,从创建开始计时,超过指定时长则
	//关闭连接，默认为0,即不关闭存活时长较长的连接
	//命令执行失败时的重试策略
	MaxRetries      int           // 命令执行失败时,最多重试多少次,默认为0即不重试
	MinRetryBackoff time.Duration //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
	MaxRetryBackoff time.Duration //每次计算重试间隔时间的上限,默认512毫秒,-1表示取消间隔
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

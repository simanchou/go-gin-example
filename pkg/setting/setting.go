package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
	"github.com/simanchou/go-gin-example/pkg/logging"
)

//App app secction config item
type App struct {
	JwtSecret string
	PageSize  int
	PrefixURL string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSavename string
	LogFileExt  string
	TimeFormat  string
}

//AppSetting app setting instance
var AppSetting = &App{}

//Server server secction config item
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//ServerSetting server setting instance
var ServerSetting = &Server{}

//Database database secction config item
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

//DatabaseSetting database setting instance
var DatabaseSetting = &Database{}

//Redis redis secction config item
type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

//RedisSetting redis setting instance
var RedisSetting = &Redis{}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v\n", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logging.Fatal(section, err)
	}
}

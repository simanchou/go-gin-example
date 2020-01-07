package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	// Cfg conf var
	Cfg *ini.File

	// RunMode run mode
	RunMode string

	// HTTPPort HTTP port
	HTTPPort int

	// ReadTimeout read timeout
	ReadTimeout time.Duration
	// WriteTimeout write timeout
	WriteTimeout time.Duration

	// PageSize page size
	PageSize int
	// JwtSecret jwt secret
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v\n", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

// LoadBase load base
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// LoadServer load server
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get setion 'server': %v\n", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIME").MustInt(60)) * time.Second
}

// LoadApp load app
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v\n", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@#$%^!@#$%^125sdfg2")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

package boot

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhouhp1295/g3-cms/utils"
	"gopkg.in/ini.v1"
	"strings"
	"unknwon.dev/clog/v2"
)

var (
	App struct {
		Product     string
		ProductName string
		Version     string
		RunMode     string
	}
)

// DatabaseConfig 数据库设置
type DatabaseConfig struct {
	Type         string
	Host         string
	Name         string
	User         string
	Password     string
	SSLMode      string `ini:"SSL_MODE"`
	Path         string
	MaxOpenConns int
	MaxIdleConns int
}

// DatabaseCfg 数据库设置
var DatabaseCfg DatabaseConfig

type loggerConfig struct {
	RootPath     string
	LogRotate    bool
	DailyRotate  bool
	MaxSizeShift int64
	MaxLines     int64
	MaxDays      int64
}

func (c *loggerConfig) ToFileRotationConfig() clog.FileRotationConfig {
	return clog.FileRotationConfig{
		Rotate:   c.LogRotate,
		Daily:    c.DailyRotate,
		MaxSize:  c.MaxSizeShift * 1024 * 1024,
		MaxDays:  c.MaxDays,
		MaxLines: c.MaxLines,
	}
}

var LoggerCfg loggerConfig

type serverConfig struct {
	Domain   string
	HTTPAddr string `ini:"HTTP_ADDR"`
	HTTPPort string `ini:"HTTP_PORT"`
}

// ServerCfg settings
var ServerCfg serverConfig

type jwtConfig struct {
	Secret         string
	ExpiredSeconds int64
}

var JwtCfg jwtConfig

type storageConfig struct {
	Type      string
	Path      string
	MaxSize   float32
	Uri       string
	UriPrefix string
}

func (cfg *storageConfig) check() bool {
	if cfg.Type == "fs" {
		if len(cfg.Path) == 0 {
			cfg.Path = HomeDir()
		}
	}
	if len(cfg.Path) == 0 {
		return false
	}
	switch cfg.Type {
	case "fs":
		if utils.IsWin() {
			cfg.Uri = fmt.Sprintf("fs://%s", strings.ReplaceAll(cfg.Path, `\`, `/`))
		} else {
			cfg.Uri = fmt.Sprintf("fs://%s", cfg.Path)
		}
	default:
		cfg.Uri = fmt.Sprintf("%s://%s", cfg.Type, cfg.Path)
	}
	if cfg.MaxSize <= 0 {
		cfg.MaxSize = 2
	}
	return true
}

var StorageCfg storageConfig

func loadConfigs() {
	var err error
	var iniPath string
	if IsInstalled() {
		iniPath = AssetPath("conf/app.ini")
	} else {
		iniPath = AssetPath("conf/app.example.ini")
	}
	if utils.IsExist(iniPath) {
		File, err = ini.LoadSources(ini.LoadOptions{
			IgnoreInlineComment: true,
		}, iniPath)
	} else {
		panic("未找到配置文件: " + iniPath)
	}

	if err != nil {
		panic(errors.Wrap(err, "配置文件解析失败: "+iniPath))
	}

	File.NameMapper = ini.SnackCase

	// *****************************
	// ----- DatabaseCfg settings -----
	// *****************************

	if err = File.Section("database").MapTo(&DatabaseCfg); err != nil {
		panic(errors.Wrap(err, "配置解析失败: database"))
	}

	// ***************************
	// ----- ServerCfg settings -----
	// ***************************
	if err = File.Section("server").MapTo(&ServerCfg); err != nil {
		panic(errors.Wrap(err, "配置解析失败: server"))
	}

	// ***************************
	// ----- JwtCfg settings -----
	// ***************************
	if err = File.Section("jwt").MapTo(&JwtCfg); err != nil {
		panic(errors.Wrap(err, "配置解析失败: jwt"))
	}

	// ***************************
	// ----- LoggerCfg settings -----
	// ***************************
	if err = File.Section("logger").MapTo(&LoggerCfg); err != nil {
		panic(errors.Wrap(err, "配置解析失败: logger"))
	}

	// ***************************
	// ----- StorageCfg settings -----
	// ***************************
	if err = File.Section("storage").MapTo(&StorageCfg); err != nil {
		panic(errors.Wrap(err, "配置解析失败: storage"))
	}
	if !StorageCfg.check() {
		panic("请检查storage配置")
	}
}

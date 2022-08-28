// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package boot

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/utils"
	"gopkg.in/ini.v1"
	"strings"
)

var File *ini.File

// DatabaseConfig 数据库设置
type DatabaseConfig struct {
	Type         string
	Host         string
	Name         string
	User         string
	Password     string
	SSLMode      string `ini:"SSL_MODE"`
	MaxOpenConns int
	MaxIdleConns int
}

// DatabaseCfg 数据库设置
var DatabaseCfg DatabaseConfig

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
			cfg.Path = g3.HomeDir()
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
		iniPath = g3.AssetPath("conf/app.ini")
	} else {
		iniPath = g3.AssetPath("conf/app.example.ini")
	}
	if !utils.IsExist(iniPath) {
		panic("未找到配置文件: " + iniPath)
	}

	File, err = ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, iniPath)

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
	// ----- JwtCfg settings -----
	// ***************************
	if err = File.Section("jwt").MapTo(&JwtCfg); err != nil {
		panic(errors.Wrap(err, "配置解析失败: jwt"))
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

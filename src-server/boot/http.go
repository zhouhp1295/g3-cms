// Copyright (c) 554949297@qq.com . 2022-2022. All rights reserved

//go:build http
// +build http

package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/utils"
	"github.com/zhouhp1295/g3/net"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
	"net/http"
	"os/exec"
	"strings"
)

func init() {
	App.Name = "http"
	App.Identifier = "node01"
	App.Version = "v0.0.1"
	App.RunMode = "dev"
	preStart = preHttpStart
	start = startHttp
}

func loadHttpConfig() {
	iniPath := g3.AssetPath("conf/http.ini")
	if !utils.IsExist(iniPath) {
		panic("未找到配置文件: " + iniPath)
	}
	iniFile, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, iniPath)

	if err != nil {
		panic(errors.Wrap(err, "配置文件解析失败: "+iniPath))
	}
	// ***************************
	// ----- ServerCfg settings -----
	// ***************************
	if err = iniFile.Section("server").MapTo(&ServerCfg); err != nil {
		panic(errors.Wrap(err, "配置解析失败: server"))
	}
}

func checkInstall(context *gin.Context) {
	if !IsInstalled() {
		if strings.HasPrefix(context.Request.RequestURI, "/install") ||
			strings.HasPrefix(context.Request.RequestURI, "/api/common") {
			context.Next()
			return
		}
		context.Redirect(http.StatusSeeOther, "/install")
		return
	}
	context.Next()
}

func preHttpStart() {
	// 加载配置
	loadHttpConfig()
	// 初始化Gin
	r := gin.Default()
	r.Use(net.DefaultCors())
	g3.SetGin(r)
	// 安装引导
	g3.GetGin().Engine.Use(checkInstall)
	g3.GetGin().Group("/api").NewJwt(JwtCfg.Secret, JwtCfg.ExpiredSeconds)
	// 初始化
	if IsInstalled() {
		DoAfterInstall()
	}
	// 打开网站
	if utils.IsMac() {
		_ = exec.Command("/bin/sh", "-c", `open http://127.0.0.1:`+ServerCfg.HTTPPort).Start()
	} else if utils.IsWin() {
		_ = exec.Command("cmd", "/c", `start http://127.0.0.1:`+ServerCfg.HTTPPort).Start()
	}
}

func startHttp() {
	err := g3.GetGin().Engine.Run(ServerCfg.HTTPAddr + ":" + ServerCfg.HTTPPort)
	if err != nil {
		g3.ZL().Fatal("Application Stopped", zap.Error(err))
	}
}

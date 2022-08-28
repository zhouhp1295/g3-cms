// Copyright (c) 554949297@qq.com . 2022-2022. All rights reserved

//go:build websocket
// +build websocket

package boot

import (
	"github.com/pkg/errors"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/utils"
	"github.com/zhouhp1295/g3/net"
	"go.uber.org/zap"
	"net/http"
)

var (
	worker *net.WsWorker
)

func init() {
	App.Name = "websocket"
	App.Identifier = "node01"
	App.Version = "v0.0.1"
	App.RunMode = "dev"
	preStart = preWebsocketStart
	start = startWebsocket
}

func loadWebsocketConfig() {
	iniPath := g3.AssetPath("conf/websocket.ini")
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

func preWebsocketStart() {
	if !IsInstalled() {
		panic("系统未安装")
	}
	// 加载配置
	loadWebsocketConfig()
	// 启动
	var err error
	worker, err = net.HandleWebsocket("/")
	if err != nil {
		g3.ZL().Fatal("服务启动失败", zap.Error(err))
	}
}

func startWebsocket() {
	err := http.ListenAndServe(ServerCfg.HTTPAddr+":"+ServerCfg.HTTPPort, nil)
	if err != nil {
		g3.ZL().Fatal("Application Stopped", zap.Error(err))
	}
}

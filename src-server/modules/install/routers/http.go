// Copyright (c) 554949297@qq.com . 2022-2022. All rights reserved

//go:build http
// +build http

package routers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/utils"
	"github.com/zhouhp1295/g3/net"
	"net/http"
	"os"
	"path"
	"sync"
)

func init() {
	boot.RegisterPreFunction(func() {
		g3.GetGin().Engine.Static("/install", path.Join(g3.HomeDir(), "public", "install"))

		g3.GetGin().Group("/api").MakeOpen("/common/install", "/common/testConn")
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/common/install", handleInstall)
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/common/testConn", handleTestConn)
	})
}

type formData struct {
	DatabaseType string `json:"databaseType"`
	DatabaseHost string `json:"databaseHost"`
	DatabaseUser string `json:"databaseUser"`
	DatabasePwd  string `json:"databasePwd"`
	DatabaseName string `json:"databaseName"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

var installMutex sync.Mutex

func handleInstall(ctx *gin.Context) {
	installMutex.Lock()
	defer installMutex.Unlock()

	if boot.IsInstalled() {
		net.FailedMessage(ctx, "已安装")
		return
	}
	var form formData
	var err error
	if err = ctx.ShouldBindJSON(&form); err != nil {
		net.FailedMessage(ctx, "参数错误")
		return
	}
	if len(form.DatabaseType) == 0 || len(form.DatabaseName) == 0 {
		net.FailedMessage(ctx, "参数错误")
		return
	}
	if len(form.Username) < 4 || len(form.Username) > 20 {
		net.FailedMessage(ctx, "请输入 4~20 长度的用户名")
		return
	}
	if len(form.Password) < 6 || len(form.Password) > 30 {
		net.FailedMessage(ctx, "请输入 4~30 长度的密码")
		return
	}
	//测试数据库连接
	ok, msg := testConn(form)
	if !ok {
		net.FailedMessage(ctx, "数据库连接失败:"+msg)
		return
	}
	//回写 cfg
	boot.File.Section("database").Key("TYPE").SetValue(form.DatabaseType)
	boot.File.Section("database").Key("HOST").SetValue(form.DatabaseHost)
	boot.File.Section("database").Key("NAME").SetValue(form.DatabaseName)
	boot.File.Section("database").Key("USER").SetValue(form.DatabaseUser)
	boot.File.Section("database").Key("PASSWORD").SetValue(form.DatabasePwd)
	boot.File.Section("database").Key("PATH").SetValue("")

	confDir := path.Join(g3.HomeDir(), "conf")
	if !utils.IsExist(confDir) {
		err := os.MkdirAll(confDir, os.ModePerm)
		if err != nil {
			net.FailedMessage(ctx, "保存配置文件失败: err = "+err.Error())
			return
		}
	}
	err = boot.File.SaveTo(path.Join(confDir, "app.ini"))
	if err != nil {
		net.FailedMessage(ctx, "保存配置文件失败: err = "+err.Error())
		return
	}
	//重新载入CFG
	if err = boot.File.Section("database").MapTo(&boot.DatabaseCfg); err != nil {
		net.FailedMessage(ctx, "保存配置文件失败: err = "+err.Error())
		return
	}

	boot.DoAfterInstall()

	net.SuccessDefault(ctx)
}

func testConn(form formData) (bool, string) {
	cfg := boot.DatabaseConfig{
		Type:     form.DatabaseType,
		Host:     form.DatabaseHost,
		User:     form.DatabaseUser,
		Password: form.DatabasePwd,
		Name:     form.DatabaseName,
	}
	if cfg.Type == boot.SQLite3 {
		dbpath := g3.AssetPath(cfg.Name + ".db")

		defer func(name string) {
			_ = os.Remove(name)
		}(dbpath)

		if utils.IsExist(dbpath) {
			return false, "数据库文件已存在"
		}
	}
	db, err := boot.TestDatabaseConn(cfg)
	if err != nil {
		return false, err.Error()
	}
	sqlDb, err := db.DB()

	defer func(sqlDb *sql.DB) {
		if sqlDb != nil {
			_ = sqlDb.Close()
		}
	}(sqlDb)

	if err != nil {
		return false, err.Error()
	}

	return true, "连接成功"
}

func handleTestConn(ctx *gin.Context) {
	if boot.IsInstalled() {
		net.FailedMessage(ctx, "已安装")
		return
	}
	var form formData
	var err error
	if err = ctx.ShouldBindJSON(&form); err != nil {
		net.FailedMessage(ctx, "参数错误")
		return
	}
	if len(form.DatabaseType) == 0 || len(form.DatabaseName) == 0 {
		net.FailedMessage(ctx, "参数错误")
		return
	}
	//测试数据库连接
	ok, msg := testConn(form)
	if ok {
		net.SuccessDefault(ctx)
	} else {
		net.FailedMessage(ctx, "数据库连接失败:"+msg)
	}
}

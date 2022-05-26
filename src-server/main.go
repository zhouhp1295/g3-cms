package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/install"
	"github.com/zhouhp1295/g3-cms/modules"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3-cms/utils"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

var showInstallGuide = true

func main() {
	//时区设置
	err := os.Setenv("TZ", "Asia/Shanghai")
	if err != nil {
		panic("初始化时区失败:" + err.Error())
		return
	}
	if showInstallGuide && !boot.IsInstalled() {
		boot.GinCtx().Engine.Use(func(context *gin.Context) {
			//初始化检测
			if !boot.IsInstalled() {
				if strings.HasPrefix(context.Request.RequestURI, "/install") ||
					strings.HasPrefix(context.Request.RequestURI, "/api/common") {
					context.Next()
					return
				}
				context.Redirect(http.StatusSeeOther, "/install")
				return
			} else {
				if strings.HasPrefix(context.Request.RequestURI, "/install") ||
					strings.HasPrefix(context.Request.RequestURI, "/api/common/install") {
					context.Redirect(http.StatusSeeOther, "/")
					return
				}
			}
			context.Next()
		})

		boot.GinCtx().Engine.Static("/install", path.Join(boot.HomeDir(), "public", "install"))

		install.InitInstallApiRouters()
	} else if !boot.IsInstalled() {
		//初始化数据库
		boot.InitDatabase()
		//同步表结构
		modules.InitModulesDatabase()
		//写入installed文件
		boot.SetInstalled()
		//创建超级管理员
		dao.SysUserDao.CreateSuperUser("admin", "123456")
		//初始化权限
		dao.SysRoleDao.RefreshRolePerms()
	}
	//绑定静态文件
	boot.BindWebStatic()

	if utils.IsMac() {
		_ = exec.Command("/bin/sh", "-c", `open http://127.0.0.1:`+boot.ServerCfg.HTTPPort).Start()
	} else if utils.IsWin() {
		_ = exec.Command("cmd", "/c", `start http://127.0.0.1:`+boot.ServerCfg.HTTPPort).Start()
	}
	boot.Logger.Info("Application Start Success , Listen Port: %s", boot.ServerCfg.HTTPPort)
	err = boot.GinCtx().Engine.Run(boot.ServerCfg.HTTPAddr + ":" + boot.ServerCfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}

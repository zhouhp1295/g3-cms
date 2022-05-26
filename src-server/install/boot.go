package install

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3-cms/utils"
	"github.com/zhouhp1295/g3/net"
	"os"
	"path"
	"sync"
)

func InitInstallApiRouters() {
	boot.SetApiOpen("/common/install", "/common/testConn")
	boot.ApiPost("/common/install", handleInstall)
	boot.ApiPost("/common/testConn", handleTestConn)
}

var api = net.BaseApi{}

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
		api.FailedMessage(ctx, "已安装")
		return
	}
	var form formData
	var err error
	if err = ctx.ShouldBindJSON(&form); err != nil {
		api.FailedMessage(ctx, "参数错误")
		return
	}
	if len(form.DatabaseType) == 0 || len(form.DatabaseName) == 0 {
		api.FailedMessage(ctx, "参数错误")
		return
	}
	if len(form.Username) < 4 || len(form.Username) > 20 {
		api.FailedMessage(ctx, "请输入 4~20 长度的用户名")
		return
	}
	if len(form.Password) < 6 || len(form.Password) > 30 {
		api.FailedMessage(ctx, "请输入 4~30 长度的密码")
		return
	}
	//测试数据库连接
	ok, msg := testConn(form)
	if !ok {
		api.FailedMessage(ctx, "数据库连接失败:"+msg)
		return
	}
	//回写 cfg
	boot.File.Section("database").Key("TYPE").SetValue(form.DatabaseType)
	boot.File.Section("database").Key("HOST").SetValue(form.DatabaseHost)
	boot.File.Section("database").Key("NAME").SetValue(form.DatabaseName)
	boot.File.Section("database").Key("USER").SetValue(form.DatabaseUser)
	boot.File.Section("database").Key("PASSWORD").SetValue(form.DatabasePwd)
	boot.File.Section("database").Key("PATH").SetValue("")

	confDir := path.Join(boot.HomeDir(), "conf")
	if !utils.IsExist(confDir) {
		err := os.MkdirAll(confDir, os.ModePerm)
		if err != nil {
			api.FailedMessage(ctx, "保存配置文件失败: err = "+err.Error())
			return
		}
	}
	err = boot.File.SaveTo(path.Join(confDir, "app.ini"))
	if err != nil {
		api.FailedMessage(ctx, "保存配置文件失败: err = "+err.Error())
		return
	}
	//重新载入CFG
	if err = boot.File.Section("database").MapTo(&boot.DatabaseCfg); err != nil {
		api.FailedMessage(ctx, "保存配置文件失败: err = "+err.Error())
		return
	}
	//初始化数据库
	boot.InitDatabase()
	//同步表结构
	modules.InitModulesDatabase()
	//创建超级管理员
	if !dao.SysUserDao.CreateSuperUser(form.Username, form.Password) {
		api.FailedMessage(ctx, "创建管理员失败，系统或数据库已初始化！")
		return
	}
	//写入installed文件
	boot.SetInstalled()
	//初始化权限
	dao.SysRoleDao.RefreshRolePerms()

	api.SuccessDefault(ctx)
}

func testConn(form formData) (bool, string) {
	cfg := boot.DatabaseConfig{
		Type:     form.DatabaseType,
		Host:     form.DatabaseHost,
		User:     form.DatabaseUser,
		Password: form.DatabasePwd,
		Name:     form.DatabaseName,
	}
	db, err := boot.TestDatabaseConn(cfg)
	if err != nil {
		return false, err.Error()
	}
	sqlDb, err := db.DB()
	if err == nil {
		_ = sqlDb.Close()
	}
	return true, "连接成功"
}

func handleTestConn(ctx *gin.Context) {
	if boot.IsInstalled() {
		api.FailedMessage(ctx, "已安装")
		return
	}
	var form formData
	var err error
	if err = ctx.ShouldBindJSON(&form); err != nil {
		api.FailedMessage(ctx, "参数错误")
		return
	}
	if len(form.DatabaseType) == 0 || len(form.DatabaseName) == 0 {
		api.FailedMessage(ctx, "参数错误")
		return
	}
	//测试数据库连接
	ok, msg := testConn(form)
	if ok {
		api.SuccessDefault(ctx)
	} else {
		api.FailedMessage(ctx, "数据库连接失败:"+msg)
	}
}

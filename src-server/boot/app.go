// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package boot

import (
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/utils"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/lache"
	"github.com/zhouhp1295/lache/driver"
	_ "go.beyondstorage.io/services/fs/v4"
	_ "go.beyondstorage.io/services/minio"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
	"go.uber.org/zap"
	"os"
	"strings"
	"sync"
)

var (
	App struct {
		Name       string
		Identifier string
		Version    string
		RunMode    string
	}
	Lache    *lache.Client
	Storager types.Storager

	initDatabaseOnce sync.Once

	installed     string
	installedOnce sync.Once

	preFunctionMutex sync.Mutex
	preFunctions     []func()

	afterInstallFunctions     []func()
	afterInstallFunctionMutex sync.Mutex

	preStart func()
	start    func()
)

func init() {

}
func RegisterPreFunction(f func()) {
	preFunctionMutex.Lock()
	preFunctions = append(preFunctions, f)
	preFunctionMutex.Unlock()
}

func RegisterAfterInstallFunction(f func()) {
	afterInstallFunctionMutex.Lock()
	afterInstallFunctions = append(afterInstallFunctions, f)
	afterInstallFunctionMutex.Unlock()
}

func DoAfterInstall() {
	//写入installed文件
	SetInstalled()
	//初始化数据库
	InitDatabase()

	for _, f := range afterInstallFunctions {
		f()
	}
}

// IsProdMode returns true if the application is running in production mode.
func IsProdMode() bool {
	return strings.EqualFold(App.RunMode, "prod")
}

// IsInstalled 是否已初始化安装
func IsInstalled() bool {
	installedOnce.Do(func() {
		if utils.IsExist(g3.AssetPath(".install")) {
			installed = crud.FlagTrue
		} else {
			installed = crud.FlagFalse
		}
	})

	if installed == crud.FlagTrue {
		return true
	}
	return false
}

func SetInstalled() {
	if installed != crud.FlagTrue {
		f, e := os.Create(g3.AssetPath(".install"))
		if e != nil {
			return
		}
		_ = f.Close()
	}
	installed = crud.FlagTrue
}

func InitDatabase() {
	initDatabaseOnce.Do(func() {
		initGormDB(&GormLogger{})
	})
}

func Start() {
	// 时区设置
	err := os.Setenv("TZ", "Asia/Shanghai")
	if err != nil {
		panic("初始化时区失败:" + err.Error())
	}
	// g3配置信息
	g3Cfg := g3.Cfg{
		HomeDir: os.Getenv("G3_CMS_HOME"),
		AppName: App.Name,
		AppId:   App.Identifier,
	}
	g3.Boot(&g3Cfg)
	// 缓存
	Lache = lache.New(lache.Local, driver.LocalOptions{})
	// 加载配置文件
	loadConfigs()
	// 存储
	Storager, err = services.NewStoragerFromString(StorageCfg.Uri)
	if err != nil {
		panic(err)
	}
	g3.ZL().Info("存储空间", zap.Reflect("Storager", Storager))
	// 运行 preStart
	if preStart != nil {
		preStart()
	}
	for _, f := range preFunctions {
		f()
	}
	g3.ZL().Info("Application Start",
		zap.String("Addr", ServerCfg.HTTPAddr),
		zap.String("Port", ServerCfg.HTTPPort))
	start()
}

package boot

import (
	"github.com/zhouhp1295/g3-cms/utils"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/lache"
	"github.com/zhouhp1295/lache/driver"
	_ "go.beyondstorage.io/services/fs/v4"
	_ "go.beyondstorage.io/services/minio"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
	"os"
	"path/filepath"
	"sync"
)

var Lache *lache.Client

var Storager types.Storager

func init() {
	//初始化配置文件
	loadConfigs()
	//初始化日志配置
	initLogger()
	//存储
	var err error
	Storager, err = services.NewStoragerFromString(StorageCfg.Uri)
	if err != nil {
		panic(err)
	}
	Logger.Info("初始化存储空间: %s", Storager.String())
	//缓存
	Lache = lache.New(lache.Local, driver.LocalOptions{})
}

var initDatabaseOnce sync.Once
var installed string

var installedOnce sync.Once

// IsInstalled 是否已初始化安装
func IsInstalled() bool {
	installedOnce.Do(func() {
		if utils.IsExist(AssetPath(".install")) {
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
		f, e := os.Create(AssetPath(".install"))
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

var (
	homeDir     string
	homeDirOnce sync.Once
)

// HomeDir 工作目录
func HomeDir() string {
	homeDirOnce.Do(func() {
		homeDir = os.Getenv("g3_cms_dist")
		if homeDir != "" {
			return
		}
		homeDir = filepath.Dir(utils.AppPath())
	})
	return homeDir
}

// EnsureAbs prepends the HomeDir to the given path if it is not an absolute path.
func EnsureAbs(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(HomeDir(), path)
}

func AssetPath(filename string) string {
	return filepath.Join(HomeDir(), filename)
}

package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var (
	appPath     string
	appPathOnce sync.Once
)

// AppPath app路径
func AppPath() string {
	appPathOnce.Do(func() {
		var err error
		appPath, err = exec.LookPath(os.Args[0])
		if err != nil {
			panic("look executable path: " + err.Error())
		}

		appPath, err = filepath.Abs(appPath)
		if err != nil {
			panic("get absolute executable path: " + err.Error())
		}
	})
	return appPath
}

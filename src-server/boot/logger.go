package boot

import (
	"github.com/pkg/errors"
	"io"
	"path/filepath"
	"unknwon.dev/clog/v2"
)

var (
	ginLoggerWriter io.Writer
)

var Logger *bootLogger

func initLogger() {
	if len(LoggerCfg.RootPath) == 0 {
		LoggerCfg.RootPath = filepath.Join(HomeDir(), "log")
	}
	var err error
	LoggerCfg.RootPath = EnsureAbs(LoggerCfg.RootPath)
	{
		ginLoggerWriter, err = clog.NewFileWriter(
			filepath.Join(LoggerCfg.RootPath, "gin.log"),
			LoggerCfg.ToFileRotationConfig(),
		)
		if err != nil {
			panic(errors.Wrap(err, `初始化日志文件失败: "gin.log"`))
			return
		}
	}
	{
		err = clog.NewFileWithName("db",
			clog.FileConfig{
				Level:              clog.LevelInfo,
				Filename:           filepath.Join(LoggerCfg.RootPath, "db.log"),
				FileRotationConfig: LoggerCfg.ToFileRotationConfig(),
			},
		)
		if err != nil {
			panic(errors.Wrap(err, `初始化日志文件失败: "db.log"`))
			return
		}
	}
	//初始化 logger
	{
		err = clog.NewFileWithName("app",
			clog.FileConfig{
				Level:              clog.LevelInfo,
				Filename:           filepath.Join(LoggerCfg.RootPath, "app.log"),
				FileRotationConfig: LoggerCfg.ToFileRotationConfig(),
			},
		)
		if err != nil {
			panic(errors.Wrap(err, `初始化日志文件失败: "app.log"`))
			return
		}
		Logger = &bootLogger{
			name: "app",
		}
	}
}

type bootLogger struct {
	name string
}

func (l *bootLogger) Debug(format string, v ...interface{}) {
	clog.TraceTo(l.name, format, v...)
}

func (l *bootLogger) Info(format string, v ...interface{}) {
	clog.InfoTo(l.name, format, v...)
}

func (l *bootLogger) Warn(format string, v ...interface{}) {
	clog.WarnTo(l.name, format, v...)
}

func (l *bootLogger) Error(format string, v ...interface{}) {
	clog.ErrorTo(l.name, format, v...)
}

func (l *bootLogger) Fatal(format string, v ...interface{}) {
	clog.FatalTo(l.name, format, v...)
}

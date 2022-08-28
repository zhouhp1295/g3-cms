// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

//go:build http
// +build http

package routers

import (
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	_ "github.com/zhouhp1295/g3-cms/modules/system/routers/http"
	"path"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Engine.Static("/backend", path.Join(g3.HomeDir(), "public", "backend"))
		g3.GetGin().Engine.Static("/upload", path.Join(g3.HomeDir(), "upload"))
		g3.GetGin().Engine.StaticFile("/favicon.ico", path.Join(g3.HomeDir(), "public", "favicon.ico"))
	})
}

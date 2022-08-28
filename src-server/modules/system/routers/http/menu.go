// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3/net"
	"net/http"
)

type _sysMenuApi struct {
	net.BaseApi
}

var SysMenuApi = &_sysMenuApi{
	net.BaseApi{Dao: dao.SysMenuDao},
}

func init() {
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/system/menu/treeOptions", SysMenuApi.treeOptions)
	})
}

func (api *_sysMenuApi) treeOptions(ctx *gin.Context) {
	net.SuccessList(ctx, dao.SysMenuDao.MenuTree())
}

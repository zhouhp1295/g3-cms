// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

//go:build http
// +build http

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3/net"
	"net/http"
)

type _dictDataApi struct {
	net.BaseApi
}

var SysDictDataApi = &_dictDataApi{
	net.BaseApi{Dao: dao.SysDictDataDao},
}

func init() {
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/system/dict/data/listByCode", SysDictDataApi.handleListByCode)
	})
}

func (api *_dictDataApi) handleListByCode(ctx *gin.Context) {
	code := ctx.Query("code")
	dataRows := dao.SysDictDataDao.FindListByColumn("code", code)
	net.SuccessList(ctx, dataRows)
}

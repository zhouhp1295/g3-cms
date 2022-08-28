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
	"go.uber.org/zap"
	"net/http"
)

type _sysConfigApi struct {
	net.BaseApi
}

var SysConfigApi = &_sysConfigApi{
	net.BaseApi{Dao: dao.SysConfigDao},
}

const (
	PermSystemConfigList   = "system:config:list"
	PermSystemConfigQuery  = "system:config:query"
	PermSystemConfigAdd    = "system:config:add"
	PermSystemConfigEdit   = "system:config:edit"
	PermSystemConfigRemove = "system:config:remove"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/system/config/page", SysConfigApi.HandlePage, PermSystemConfigQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/system/config/get", SysConfigApi.HandleGet, PermSystemConfigQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/admin/system/config/insert", SysConfigApi.HandleInsert, PermSystemConfigAdd)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/system/config/update", SysConfigApi.HandleUpdate, PermSystemConfigEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/system/config/status", SysConfigApi.HandleUpdateStatus, PermSystemConfigEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodDelete, "/admin/system/config/delete", SysConfigApi.HandleDelete, PermSystemConfigRemove)
		//扩展
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/system/config/getByCode", SysConfigApi.getByCode)
	})
}

type getByCodeParam struct {
	Code string `json:"code" form:"code"`
}

func (api *_sysConfigApi) getByCode(ctx *gin.Context) {
	params := getByCodeParam{}
	err := net.ShouldBind(ctx, &params)
	if err != nil {
		g3.ZL().Error("parse params failed. please check")
		net.FailedMessage(ctx, "参数错误")
		return
	}
	if len(params.Code) == 0 {
		g3.ZL().Error("code is empty. please check")
		net.FailedMessage(ctx, "参数错误")
		return
	}
	cnt := dao.SysConfigDao.CountByColumn("code", params.Code)
	if cnt <= 0 {
		g3.ZL().Error("record not exist. please check", zap.String("code", params.Code))
		net.FailedNotFound(ctx)
		return
	}
	net.SuccessData(ctx, dao.SysConfigDao.FindOneByColumn("code", params.Code))
}

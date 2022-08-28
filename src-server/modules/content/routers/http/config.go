// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

//go:build http
// +build http

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3/auth"
	"github.com/zhouhp1295/g3/net"
	"go.uber.org/zap"
	"net/http"
)

const (
	PermContentConfigList  = "content:config:list"
	PermContentConfigQuery = "content:config:query"
	PermContentConfigEdit  = "content:config:edit"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/config/web/get", getWebConfigHandler, PermContentConfigQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/config/web/update", updateWebConfigHandler, PermContentConfigEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/config/cache/clean", cacheCleanHandler, PermContentConfigEdit)
	})
}

func getWebConfigHandler(ctx *gin.Context) {
	net.SuccessData(ctx, dao.ContentConfigDao.WebConfig())
}

func updateWebConfigHandler(ctx *gin.Context) {
	params := dao.ContentWebConfigData{}
	err := net.ShouldBind(ctx, &params)
	if err != nil {
		g3.ZL().Error("parse params failed. please check")
		net.FailedMessage(ctx, "参数错误")
	}
	msg, ok := dao.ContentConfigDao.UpdateWebConfig(params, ctx.GetInt64(auth.CtxJwtUid))
	if ok {
		net.SuccessDefault(ctx)
	} else {
		g3.ZL().Error("update failed. please check", zap.Reflect("data", params))
		net.FailedMessage(ctx, "操作失败,err="+msg)
	}
}

func cacheCleanHandler(ctx *gin.Context) {
	dao.ContentConfigDao.Clean()
	net.SuccessDefault(ctx)
}

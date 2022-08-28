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

type _contentMenuApi struct {
	net.BaseApi
}

var ContentMenuApi = &_contentMenuApi{
	net.BaseApi{Dao: dao.ContentMenuDao},
}

const (
	PermContentMenuList   = "content:menu:list"
	PermContentMenuQuery  = "content:menu:query"
	PermContentMenuAdd    = "content:menu:add"
	PermContentMenuEdit   = "content:menu:edit"
	PermContentMenuRemove = "content:menu:remove"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		//基础接口
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/menu/list", ContentMenuApi.HandleList, PermContentMenuQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/menu/get", ContentMenuApi.HandleGet, PermContentMenuQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/admin/content/menu/insert", ContentMenuApi.HandleInsert, PermContentMenuAdd)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/menu/update", ContentMenuApi.HandleUpdate, PermContentMenuEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/menu/status", ContentMenuApi.HandleUpdateStatus, PermContentMenuEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodDelete, "/admin/content/menu/delete", ContentMenuApi.HandleDelete, PermContentMenuRemove)
		//扩展
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/menu/isBlank", ContentMenuApi.handleUpdateIsBlank, PermContentMenuRemove)
	})
}

type updateMenuParams struct {
	net.IdParams
	IsBlank string `json:"isBlank" form:"isBlank"`
}

func (api *_contentMenuApi) handleUpdateIsBlank(ctx *gin.Context) {
	params := updateMenuParams{}
	err := net.ShouldBind(ctx, &params)
	if err != nil {
		g3.ZL().Error("parse params failed. please check")
		net.FailedMessage(ctx, "参数错误")
		return
	}
	if api.Dao.CountByPk(params.Id) == 0 {
		g3.ZL().Error("record not exist. please check", zap.Int64("id", params.Id))
		net.FailedNotFound(ctx)
		return
	}
	if len(params.IsBlank) == 0 {
		g3.ZL().Error("parse params failed. please check")
		net.FailedMessage(ctx, "参数错误")
		return
	}
	operator := ctx.GetInt64(auth.CtxJwtUid)

	if api.Dao.UpdateColumn(params.Id, "is_blank", params.IsBlank, operator) {
		net.SuccessDefault(ctx)
	} else {
		g3.ZL().Error("update failed. please check", zap.Reflect("data", params))
		net.FailedMessage(ctx, "操作失败, 请稍后重试")
	}
}

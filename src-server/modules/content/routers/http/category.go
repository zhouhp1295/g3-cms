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

type _contentCategoryApi struct {
	net.BaseApi
}

var ContentCategoryApi = &_contentCategoryApi{
	net.BaseApi{Dao: dao.ContentCategoryDao},
}

const (
	PermContentCategoryList   = "content:category:list"
	PermContentCategoryQuery  = "content:category:query"
	PermContentCategoryAdd    = "content:category:add"
	PermContentCategoryEdit   = "content:category:edit"
	PermContentCategoryRemove = "content:category:remove"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		//基础接口
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/category/list", ContentCategoryApi.HandleList, PermContentCategoryQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/category/get", ContentCategoryApi.HandleGet, PermContentCategoryQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/admin/content/category/insert", ContentCategoryApi.HandleInsert, PermContentCategoryAdd)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/category/update", ContentCategoryApi.HandleUpdate, PermContentCategoryEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/category/status", ContentCategoryApi.HandleUpdateStatus, PermContentCategoryEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodDelete, "/admin/content/category/delete", ContentCategoryApi.HandleDelete, PermContentCategoryRemove)
		//扩展
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/category/inMenu", ContentCategoryApi.handleUpdateInMenu, PermContentCategoryEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/category/inBanner", ContentCategoryApi.handleUpdateInBanner, PermContentCategoryEdit)
	})
}

func (api *_contentCategoryApi) handleUpdateInMenu(ctx *gin.Context) {
	params := updateMenuOrBannerParams{}
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
	operator := ctx.GetInt64(auth.CtxJwtUid)

	if dao.ContentCategoryDao.UpdateInMenu(params.Id, params.InMenu, params.InMenuSort, operator) {
		net.SuccessDefault(ctx)
		return
	}
	g3.ZL().Error("update failed. please check", zap.Reflect("data", params))
	net.FailedMessage(ctx, "操作失败, 请稍后重试.")
}

func (api *_contentCategoryApi) handleUpdateInBanner(ctx *gin.Context) {
	params := updateMenuOrBannerParams{}
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
	operator := ctx.GetInt64(auth.CtxJwtUid)

	if dao.ContentCategoryDao.UpdateInBanner(params.Id, params.InBanner, params.InBannerSort, operator) {
		net.SuccessDefault(ctx)
		return
	}
	g3.ZL().Error("update failed. please check", zap.Reflect("data", params))
	net.FailedMessage(ctx, "操作失败, 请稍后重试.")
}

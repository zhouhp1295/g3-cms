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

type _contentArticleApi struct {
	net.BaseApi
}

var contentArticleApi = &_contentArticleApi{
	net.BaseApi{Dao: dao.ContentArticleDao},
}

const (
	PermContentArticleList   = "content:article:list"
	PermContentArticleQuery  = "content:article:query"
	PermContentArticleAdd    = "content:article:add"
	PermContentArticleEdit   = "content:article:edit"
	PermContentArticleRemove = "content:article:remove"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		//基础接口
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/article/page", contentArticleApi.HandlePage, PermContentArticleQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/article/get", contentArticleApi.HandleGet, PermContentArticleQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/admin/content/article/insert", contentArticleApi.HandleInsert, PermContentArticleAdd)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/article/update", contentArticleApi.HandleUpdate, PermContentArticleEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/article/status", contentArticleApi.HandleUpdateStatus, PermContentArticleEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodDelete, "/admin/content/article/delete", contentArticleApi.HandleDelete, PermContentArticleRemove)
		//扩展
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/article/inBanner", contentArticleApi.handleUpdateInBanner, PermContentArticleEdit)
	})
}

type updateMenuOrBannerParams struct {
	net.IdParams
	InMenu       string `json:"inMenu" form:"inMenu"`
	InMenuSort   int    `json:"inMenuSort" form:"inMenuSort"`
	InBanner     string `json:"inBanner" form:"inBanner"`
	InBannerSort int    `json:"inBannerSort" form:"inBannerSort"`
}

func (api *_contentArticleApi) handleUpdateInBanner(ctx *gin.Context) {
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

	if dao.ContentArticleDao.UpdateInBanner(params.Id, params.InBanner, params.InBannerSort, operator) {
		net.SuccessDefault(ctx)
		return
	}
	g3.ZL().Error("update failed. please check", zap.Reflect("data", params))
	net.FailedMessage(ctx, "操作失败, 请稍后重试")
}

// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

//go:build http
// +build http

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3/net"
	"go.uber.org/zap"
	"net/http"
)

type _contentTagApi struct {
	net.BaseApi
}

var ContentTagApi = &_contentTagApi{
	net.BaseApi{Dao: dao.ContentTagDao},
}

const (
	PermContentTagList   = "content:tag:list"
	PermContentTagQuery  = "content:tag:query"
	PermContentTagAdd    = "content:tag:add"
	PermContentTagEdit   = "content:tag:edit"
	PermContentTagRemove = "content:tag:remove"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		//基础接口
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/tag/page", ContentTagApi.HandlePage, PermContentTagQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/tag/get", ContentTagApi.HandleGet, PermContentTagQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/admin/content/tag/insert", ContentTagApi.HandleInsert, PermContentTagAdd)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/tag/update", ContentTagApi.HandleUpdate, PermContentTagEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/tag/status", ContentTagApi.HandleUpdateStatus, PermContentTagEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodDelete, "/admin/content/tag/delete", ContentTagApi.HandleDelete, PermContentTagRemove)
		//扩展接口
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/admin/content/tag/fastInsert", ContentTagApi.handleFastInsert)
	})
}

type tagFastInsertParams struct {
	Title string `json:"title" form:"title"`
}

func (api *_contentTagApi) handleFastInsert(ctx *gin.Context) {
	params := tagFastInsertParams{}
	err := net.ShouldBind(ctx, &params)
	if err != nil {
		g3.ZL().Error("parse params failed. please check")
		net.FailedMessage(ctx, "参数错误")
		return
	}
	m := dao.ContentTagDao.InsertOrGetByTitle(params.Title)
	if m == nil {
		g3.ZL().Error("insert or get failed. please check", zap.Reflect("data", params))
		net.FailedMessage(ctx, "操作失败,请稍后重试")
		return
	}
	net.SuccessData(ctx, gin.H{"id": m.Id, "title": m.Title})
}

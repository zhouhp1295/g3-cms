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
	"net/http"
)

type _contentWriterApi struct {
	net.BaseApi
}

var ContentWriterApi = &_contentWriterApi{
	net.BaseApi{Dao: dao.ContentWriterDao},
}

const (
	PermContentWriterList   = "content:writer:list"
	PermContentWriterQuery  = "content:writer:query"
	PermContentWriterAdd    = "content:writer:add"
	PermContentWriterEdit   = "content:writer:edit"
	PermContentWriterRemove = "content:writer:remove"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		//基础接口
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/writer/page", ContentWriterApi.HandlePage, PermContentWriterQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/writer/get", ContentWriterApi.HandleGet, PermContentWriterQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/admin/content/writer/insert", ContentWriterApi.HandleInsert, PermContentWriterAdd)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/writer/update", ContentWriterApi.HandleUpdate, PermContentWriterEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/writer/status", ContentWriterApi.HandleUpdateStatus, PermContentWriterEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodDelete, "/admin/content/writer/delete", ContentWriterApi.HandleDelete, PermContentWriterRemove)
		//扩展
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/writer/selectOptions", ContentWriterApi.selectOptions)
	})
}

func (api *_contentWriterApi) selectOptions(ctx *gin.Context) {
	net.SuccessList(ctx, dao.ContentWriterDao.SelectOptions())
}

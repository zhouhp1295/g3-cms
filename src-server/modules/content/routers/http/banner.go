// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

//go:build http
// +build http

package api

import (
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3/net"
	"net/http"
)

type _contentBannerApi struct {
	net.BaseApi
}

var ContentBannerApi = &_contentBannerApi{
	net.BaseApi{Dao: dao.ContentBannerDao},
}

const (
	PermContentBannerList   = "content:banner:list"
	PermContentBannerQuery  = "content:banner:query"
	PermContentBannerAdd    = "content:banner:add"
	PermContentBannerEdit   = "content:banner:edit"
	PermContentBannerRemove = "content:banner:remove"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		//基础接口
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/banner/page", ContentBannerApi.HandlePage, PermContentBannerQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/banner/get", ContentBannerApi.HandleGet, PermContentBannerQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/admin/content/banner/insert", ContentBannerApi.HandleInsert, PermContentBannerAdd)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/banner/update", ContentBannerApi.HandleUpdate, PermContentBannerEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/banner/status", ContentBannerApi.HandleUpdateStatus, PermContentBannerEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodDelete, "/admin/content/banner/delete", ContentBannerApi.HandleDelete, PermContentBannerRemove)
	})
}

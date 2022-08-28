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

type _contentFriendLinkApi struct {
	net.BaseApi
}

var ContentFriendLinkApi = &_contentFriendLinkApi{
	net.BaseApi{Dao: dao.ContentFriendLinkDao},
}

const (
	PermContentFriendLinkList   = "content:friendLink:list"
	PermContentFriendLinkQuery  = "content:friendLink:query"
	PermContentFriendLinkAdd    = "content:friendLink:add"
	PermContentFriendLinkEdit   = "content:friendLink:edit"
	PermContentFriendLinkRemove = "content:friendLink:remove"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		//基础接口
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/friendLink/page", ContentFriendLinkApi.HandlePage, PermContentFriendLinkQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/content/friendLink/get", ContentFriendLinkApi.HandleGet, PermContentFriendLinkQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/admin/content/friendLink/insert", ContentFriendLinkApi.HandleInsert, PermContentFriendLinkAdd)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/friendLink/update", ContentFriendLinkApi.HandleUpdate, PermContentFriendLinkEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/content/friendLink/status", ContentFriendLinkApi.HandleUpdateStatus, PermContentFriendLinkEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodDelete, "/admin/content/friendLink/delete", ContentFriendLinkApi.HandleDelete, PermContentFriendLinkRemove)
	})
}

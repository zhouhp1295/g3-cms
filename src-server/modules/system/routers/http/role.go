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

type _sysRoleApi struct {
	net.BaseApi
}

var SysRoleApi = &_sysRoleApi{
	net.BaseApi{Dao: dao.SysRoleDao},
}

const (
	PermSystemRoleList   = "system:role:list"
	PermSystemRoleQuery  = "system:role:query"
	PermSystemRoleAdd    = "system:role:add"
	PermSystemRoleEdit   = "system:role:edit"
	PermSystemRoleRemove = "system:role:remove"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/system/role/page", SysRoleApi.HandlePage, PermSystemRoleQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/system/role/get", SysRoleApi.HandleGet, PermSystemRoleQuery)
		g3.GetGin().Group("/api").
			Bind(http.MethodPost, "/admin/system/role/insert", SysRoleApi.HandleInsert, PermSystemRoleAdd)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/system/role/update", SysRoleApi.HandleUpdate, PermSystemRoleEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodPut, "/admin/system/role/status", SysRoleApi.HandleUpdateStatus, PermSystemRoleEdit)
		g3.GetGin().Group("/api").
			Bind(http.MethodDelete, "/admin/system/role/delete", SysRoleApi.HandleDelete, PermSystemRoleRemove)
		//扩展接口
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/admin/system/role/listOptions", SysRoleApi.listOptions)
	})
}

func (api *_sysRoleApi) listOptions(ctx *gin.Context) {
	net.SuccessList(ctx, dao.SysRoleDao.RoleOptions())
}

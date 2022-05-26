package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3/net"
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
	boot.ApiGet("/admin/system/role/page", SysRoleApi.HandlePage, PermSystemRoleQuery)
	boot.ApiGet("/admin/system/role/get", SysRoleApi.HandleGet, PermSystemRoleQuery)
	boot.ApiPost("/admin/system/role/insert", SysRoleApi.HandleInsert, PermSystemRoleAdd)
	boot.ApiPut("/admin/system/role/update", SysRoleApi.HandleUpdate, PermSystemRoleEdit)
	boot.ApiPut("/admin/system/role/status", SysRoleApi.HandleUpdateStatus, PermSystemRoleEdit)
	boot.ApiDelete("/admin/system/role/delete", SysRoleApi.HandleDelete, PermSystemRoleRemove)
	//扩展接口
	boot.ApiGet("/admin/system/role/listOptions", SysRoleApi.listOptions)
}

func (api *_sysRoleApi) listOptions(ctx *gin.Context) {
	api.SuccessList(ctx, dao.SysRoleDao.RoleOptions())
}

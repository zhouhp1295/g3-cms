package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3/net"
)

type _sysConfigApi struct {
	net.BaseApi
}

var SysConfigApi = &_sysConfigApi{
	net.BaseApi{Dao: dao.SysConfigDao},
}

const (
	PermSystemConfigList   = "system:config:list"
	PermSystemConfigQuery  = "system:config:query"
	PermSystemConfigAdd    = "system:config:add"
	PermSystemConfigEdit   = "system:config:edit"
	PermSystemConfigRemove = "system:config:remove"
)

func init() {
	boot.ApiGet("/admin/system/config/page", SysConfigApi.HandlePage, PermSystemConfigQuery)
	boot.ApiGet("/admin/system/config/get", SysConfigApi.HandleGet, PermSystemConfigQuery)
	boot.ApiPost("/admin/system/config/insert", SysConfigApi.HandleInsert, PermSystemConfigAdd)
	boot.ApiPut("/admin/system/config/update", SysConfigApi.HandleUpdate, PermSystemConfigEdit)
	boot.ApiPut("/admin/system/config/status", SysConfigApi.HandleUpdateStatus, PermSystemConfigEdit)
	boot.ApiDelete("/admin/system/config/delete", SysConfigApi.HandleDelete, PermSystemConfigRemove)
	//扩展
	boot.ApiGet("/admin/system/config/getByCode", SysConfigApi.getByCode)
}

type getByCodeParam struct {
	Code string `json:"code" form:"code"`
}

func (api *_sysConfigApi) getByCode(ctx *gin.Context) {
	param := getByCodeParam{}
	if ctx.ShouldBindQuery(&param) != nil {
		api.FailedMessage(ctx, "参数错误")
		return
	}
	if len(param.Code) == 0 {
		api.FailedMessage(ctx, "参数错误")
		return
	}
	cnt := dao.SysConfigDao.CountByColumn("code", param.Code)
	if cnt <= 0 {
		api.FailedNotFound(ctx)
		return
	}
	api.SuccessData(ctx, dao.SysConfigDao.FindOneByColumn("code", param.Code))
}

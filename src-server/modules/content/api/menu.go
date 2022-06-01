package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3/auth"
	"github.com/zhouhp1295/g3/net"
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
	//基础接口
	boot.ApiGet("/admin/content/menu/list", ContentMenuApi.HandleList, PermContentMenuQuery)
	boot.ApiGet("/admin/content/menu/get", ContentMenuApi.HandleGet, PermContentMenuQuery)
	boot.ApiPost("/admin/content/menu/insert", ContentMenuApi.HandleInsert, PermContentMenuAdd)
	boot.ApiPut("/admin/content/menu/update", ContentMenuApi.HandleUpdate, PermContentMenuEdit)
	boot.ApiPut("/admin/content/menu/status", ContentMenuApi.HandleUpdateStatus, PermContentMenuEdit)
	boot.ApiDelete("/admin/content/menu/delete", ContentMenuApi.HandleDelete, PermContentMenuRemove)
	//扩展
	boot.ApiPut("/admin/content/menu/isBlank", ContentMenuApi.handleUpdateIsBlank, PermContentMenuRemove)
}

type updateMenuParams struct {
	net.IdParams
	IsBlank string `json:"isBlank" form:"isBlank"`
}

func (api *_contentMenuApi) handleUpdateIsBlank(ctx *gin.Context) {
	params := updateMenuParams{}
	var err error
	switch net.Method(ctx.Request.Method) {
	case http.MethodGet:
		if err = ctx.ShouldBindQuery(&params); err != nil {
			api.FailedMessage(ctx, "参数错误")
			return
		}
	case http.MethodPost:
		fallthrough
	case http.MethodPut:
		if err = ctx.ShouldBindJSON(&params); err != nil {
			api.FailedMessage(ctx, "参数错误")
			return
		}
	default:
		api.FailedMessage(ctx, "METHOD错误")
		return
	}
	if len(params.IsBlank) == 0 {
		api.FailedMessage(ctx, "参数错误")
		return
	}
	if api.Dao.CountByPk(params.Id) == 0 {
		api.FailedNotFound(ctx)
		return
	}
	operator := ctx.GetInt64(auth.CtxJwtUid)

	if api.Dao.UpdateColumn(params.Id, "is_blank", params.IsBlank, operator) {
		api.SuccessDefault(ctx)
	} else {
		api.FailedMessage(ctx, "操作失败, 请稍后重试")
	}
}

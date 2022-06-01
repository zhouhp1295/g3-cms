package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/auth"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/net"
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
	//基础接口
	boot.ApiGet("/admin/content/category/list", ContentCategoryApi.HandleList, PermContentCategoryQuery)
	boot.ApiGet("/admin/content/category/get", ContentCategoryApi.HandleGet, PermContentCategoryQuery)
	boot.ApiPost("/admin/content/category/insert", ContentCategoryApi.HandleInsert, PermContentCategoryAdd)
	boot.ApiPut("/admin/content/category/update", ContentCategoryApi.HandleUpdate, PermContentCategoryEdit)
	boot.ApiPut("/admin/content/category/status", ContentCategoryApi.HandleUpdateStatus, PermContentCategoryEdit)
	boot.ApiDelete("/admin/content/category/delete", ContentCategoryApi.HandleDelete, PermContentCategoryRemove)
	//扩展
	boot.ApiPut("/admin/content/category/inMenu", ContentCategoryApi.handleUpdateInMenu, PermContentCategoryEdit)
	boot.ApiPut("/admin/content/category/inBanner", ContentCategoryApi.handleUpdateInBanner, PermContentCategoryEdit)
}

func (api *_contentCategoryApi) handleUpdateInMenu(ctx *gin.Context) {
	params := updateMenuOrBannerParams{}
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
	if len(params.InMenu) == 0 {
		api.FailedMessage(ctx, "参数错误")
		return
	}
	if api.Dao.CountByPk(params.Id) == 0 {
		api.FailedNotFound(ctx)
		return
	}
	operator := ctx.GetInt64(auth.CtxJwtUid)

	category := new(model.ContentCategory)
	category.Id = params.Id
	category.InMenu = params.InMenu
	category.InMenuSort = params.InMenuSort
	category.UpdatedBy = operator

	err = crud.DbSess().Select([]string{"in_banner", "in_banner_sort", "updated_by", "updated_at"}).Updates(category).Error

	if err == nil {
		api.SuccessDefault(ctx)
	} else {
		api.FailedMessage(ctx, "操作失败, 请稍后重试。err="+err.Error())
	}
}

func (api *_contentCategoryApi) handleUpdateInBanner(ctx *gin.Context) {
	params := updateMenuOrBannerParams{}
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
	if len(params.InBanner) == 0 {
		api.FailedMessage(ctx, "参数错误")
		return
	}
	if api.Dao.CountByPk(params.Id) == 0 {
		api.FailedNotFound(ctx)
		return
	}
	operator := ctx.GetInt64(auth.CtxJwtUid)

	category := new(model.ContentCategory)
	category.Id = params.Id
	category.InBanner = params.InBanner
	category.InBannerSort = params.InBannerSort
	category.UpdatedBy = operator

	err = crud.DbSess().Select([]string{"in_banner", "in_banner_sort", "updated_by", "updated_at"}).Updates(category).Error

	if err == nil {
		api.SuccessDefault(ctx)
	} else {
		api.FailedMessage(ctx, "操作失败, 请稍后重试。err="+err.Error())
	}
}

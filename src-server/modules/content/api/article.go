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

type _contentArticleApi struct {
	net.BaseApi
}

var ContentArticleApi = &_contentArticleApi{
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
	//基础接口
	boot.ApiGet("/admin/content/article/page", ContentArticleApi.HandlePage, PermContentArticleQuery)
	boot.ApiGet("/admin/content/article/get", ContentArticleApi.HandleGet, PermContentArticleQuery)
	boot.ApiPost("/admin/content/article/insert", ContentArticleApi.HandleInsert, PermContentArticleAdd)
	boot.ApiPut("/admin/content/article/update", ContentArticleApi.HandleUpdate, PermContentArticleEdit)
	boot.ApiPut("/admin/content/article/status", ContentArticleApi.HandleUpdateStatus, PermContentArticleEdit)
	boot.ApiDelete("/admin/content/article/delete", ContentArticleApi.HandleDelete, PermContentArticleRemove)
	//扩展
	boot.ApiPut("/admin/content/article/inBanner", ContentArticleApi.handleUpdateInBanner, PermContentArticleEdit)

}

func (api *_contentArticleApi) handleUpdateInBanner(ctx *gin.Context) {
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

	article := new(model.ContentArticle)
	article.Id = params.Id
	article.InBanner = params.InBanner
	article.InBannerSort = params.InBannerSort
	article.UpdatedBy = operator

	err = crud.DbSess().Select([]string{"in_banner", "in_banner_sort", "updated_by", "updated_at"}).Updates(article).Error

	if err == nil {
		api.SuccessDefault(ctx)
	} else {
		api.FailedMessage(ctx, "操作失败, 请稍后重试。err="+err.Error())
	}
}

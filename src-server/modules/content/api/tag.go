package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3/net"
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
	//基础接口
	boot.ApiGet("/admin/content/tag/page", ContentTagApi.HandlePage, PermContentTagQuery)
	boot.ApiGet("/admin/content/tag/get", ContentTagApi.HandleGet, PermContentTagQuery)
	boot.ApiPost("/admin/content/tag/insert", ContentTagApi.HandleInsert, PermContentTagAdd)
	boot.ApiPut("/admin/content/tag/update", ContentTagApi.HandleUpdate, PermContentTagEdit)
	boot.ApiPut("/admin/content/tag/status", ContentTagApi.HandleUpdateStatus, PermContentTagEdit)
	boot.ApiDelete("/admin/content/tag/delete", ContentTagApi.HandleDelete, PermContentTagRemove)
	//扩展接口
	boot.ApiPost("/admin/content/tag/fastInsert", ContentTagApi.handleFastInsert)
}

type tagFastInsertParams struct {
	Title string `json:"title" form:"title"`
}

func (api *_contentTagApi) handleFastInsert(ctx *gin.Context) {
	params := new(tagFastInsertParams)
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
	m := dao.ContentTagDao.InsertOrGetByTitle(params.Title)
	if m == nil {
		api.FailedMessage(ctx, "操作失败,请稍后重试")
		return
	}
	api.SuccessData(ctx, gin.H{"id": m.Id, "title": m.Title})
}

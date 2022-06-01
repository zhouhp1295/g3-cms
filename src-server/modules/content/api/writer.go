package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3/net"
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
	//基础接口
	boot.ApiGet("/admin/content/writer/page", ContentWriterApi.HandlePage, PermContentWriterQuery)
	boot.ApiGet("/admin/content/writer/get", ContentWriterApi.HandleGet, PermContentWriterQuery)
	boot.ApiPost("/admin/content/writer/insert", ContentWriterApi.HandleInsert, PermContentWriterAdd)
	boot.ApiPut("/admin/content/writer/update", ContentWriterApi.HandleUpdate, PermContentWriterEdit)
	boot.ApiPut("/admin/content/writer/status", ContentWriterApi.HandleUpdateStatus, PermContentWriterEdit)
	boot.ApiDelete("/admin/content/writer/delete", ContentWriterApi.HandleDelete, PermContentWriterRemove)
	//扩展
	boot.ApiGet("/admin/content/writer/selectOptions", ContentWriterApi.selectOptions)
}

func (api *_contentWriterApi) selectOptions(ctx *gin.Context) {
	api.SuccessList(ctx, dao.ContentWriterDao.SelectOptions())
}

package api

import (
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3/net"
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
	//基础接口
	boot.ApiGet("/admin/content/banner/page", ContentBannerApi.HandlePage, PermContentBannerQuery)
	boot.ApiGet("/admin/content/banner/get", ContentBannerApi.HandleGet, PermContentBannerQuery)
	boot.ApiPost("/admin/content/banner/insert", ContentBannerApi.HandleInsert, PermContentBannerAdd)
	boot.ApiPut("/admin/content/banner/update", ContentBannerApi.HandleUpdate, PermContentBannerEdit)
	boot.ApiPut("/admin/content/banner/status", ContentBannerApi.HandleUpdateStatus, PermContentBannerEdit)
	boot.ApiDelete("/admin/content/banner/delete", ContentBannerApi.HandleDelete, PermContentBannerRemove)
}

package api

import (
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3/net"
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
	//基础接口
	boot.ApiGet("/admin/content/friendLink/page", ContentFriendLinkApi.HandlePage, PermContentFriendLinkQuery)
	boot.ApiGet("/admin/content/friendLink/get", ContentFriendLinkApi.HandleGet, PermContentFriendLinkQuery)
	boot.ApiPost("/admin/content/friendLink/insert", ContentFriendLinkApi.HandleInsert, PermContentFriendLinkAdd)
	boot.ApiPut("/admin/content/friendLink/update", ContentFriendLinkApi.HandleUpdate, PermContentFriendLinkEdit)
	boot.ApiPut("/admin/content/friendLink/status", ContentFriendLinkApi.HandleUpdateStatus, PermContentFriendLinkEdit)
	boot.ApiDelete("/admin/content/friendLink/delete", ContentFriendLinkApi.HandleDelete, PermContentFriendLinkRemove)
}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3/net"
)

type _sysMenuApi struct {
	net.BaseApi
}

var SysMenuApi = &_sysMenuApi{
	net.BaseApi{Dao: dao.SysMenuDao},
}

func init() {
	boot.ApiGet("/admin/system/menu/treeOptions", SysMenuApi.treeOptions)
}

func (api *_sysMenuApi) treeOptions(ctx *gin.Context) {
	api.SuccessList(ctx, dao.SysMenuDao.MenuTree())
}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3/net"
)

type _dictDataApi struct {
	net.BaseApi
}

var SysDictDataApi = &_dictDataApi{
	net.BaseApi{Dao: dao.SysDictDataDao},
}

func init() {
	boot.ApiGet("/admin/system/dict/data/listByCode", SysDictDataApi.handleListByCode)
}

func (api *_dictDataApi) handleListByCode(ctx *gin.Context) {
	code := ctx.Query("code")
	dataRows := dao.SysDictDataDao.FindListByColumn("code", code)
	api.SuccessList(ctx, dataRows)
}

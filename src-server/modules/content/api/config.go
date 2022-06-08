package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/common/api"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3/auth"
	"github.com/zhouhp1295/g3/net"
	"net/http"
)

const (
	PermContentConfigList  = "content:config:list"
	PermContentConfigQuery = "content:config:query"
	PermContentConfigEdit  = "content:config:edit"
)

func init() {
	boot.ApiGet("/admin/content/config/web/get", getWebConfigHandler, PermContentConfigQuery)
	boot.ApiPut("/admin/content/config/web/update", updateWebConfigHandler, PermContentConfigEdit)
	boot.ApiPut("/admin/content/config/cache/clean", cacheCleanHandler, PermContentConfigEdit)
}

func getWebConfigHandler(ctx *gin.Context) {
	api.CommonApi.SuccessData(ctx, dao.ContentConfigDao.WebConfig())
}

func updateWebConfigHandler(ctx *gin.Context) {
	params := dao.ContentWebConfigData{}
	var err error
	switch net.Method(ctx.Request.Method) {
	case http.MethodGet:
		if err = ctx.ShouldBindQuery(&params); err != nil {
			api.CommonApi.FailedMessage(ctx, "参数错误")
			return
		}
	case http.MethodPost:
		fallthrough
	case http.MethodPut:
		if err = ctx.ShouldBindJSON(&params); err != nil {
			api.CommonApi.FailedMessage(ctx, "参数错误")
			return
		}
	default:
		api.CommonApi.FailedMessage(ctx, "METHOD错误")
		return
	}
	msg, ok := dao.ContentConfigDao.UpdateWebConfig(params, ctx.GetInt64(auth.CtxJwtUid))
	if ok {
		api.CommonApi.SuccessDefault(ctx)
	} else {
		api.CommonApi.FailedMessage(ctx, "操作失败,err="+msg)
	}
}

func cacheCleanHandler(ctx *gin.Context) {
	dao.ContentConfigDao.Clean()
	api.CommonApi.SuccessDefault(ctx)
}

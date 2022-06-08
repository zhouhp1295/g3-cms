package content

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3-cms/modules/content/migrations"
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3-cms/modules/content/web"
	"github.com/zhouhp1295/g3/crud"
	"net/http"
)

func DoMigrate() {
	crud.DoMigrate(migrations.M20220531InitContentCode, migrations.M20220531InitContent())
}

func SyncTables() {
	//初始化数据结构
	tables := []interface{}{
		new(model.ContentArticle),
		new(model.ContentArticleTag),
		new(model.ContentBanner),
		new(model.ContentCategory),
		new(model.ContentFriendLink),
		new(model.ContentMenu),
		new(model.ContentTag),
		new(model.ContentWriter),
	}
	err := crud.SyncTables(crud.DbSess(), tables)
	if err != nil {
		boot.Logger.Fatal("AutoMigrate Content Database, err = %s", err.Error())
	}
}

func InitWeb() {
	//页面
	boot.WebPage("/robots.txt", robotsHandler)
	//初始化 todo 根据设置的模板引擎去设置 web or api
	web.Init()
}

func robotsHandler(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(dao.ContentConfigDao.WebConfig().Robots))
}

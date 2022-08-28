// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package content

import (
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/modules/content/migrations"
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	_ "github.com/zhouhp1295/g3-cms/modules/content/routers"
	"github.com/zhouhp1295/g3/crud"
	"go.uber.org/zap"
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
		g3.ZL().Fatal("AutoMigrate Content Database", zap.Error(err))
	}
}

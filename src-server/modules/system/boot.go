// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package system

import (
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/modules/system/migrations"
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	_ "github.com/zhouhp1295/g3-cms/modules/system/routers"
	"github.com/zhouhp1295/g3/crud"
	"go.uber.org/zap"
)

func DoMigrate() {
	crud.DoMigrate(migrations.M20220519InitSystemCode, migrations.M20220519InitSystem())
}

func SyncTables() {
	//初始化数据结构
	tables := []interface{}{
		new(model.SysUser),
		new(model.SysRole),
		new(model.SysMenu),
		new(model.SysConfig),
		new(model.SysDictType),
		new(model.SysDictData),
	}
	err := crud.SyncTables(crud.DbSess(), tables)
	if err != nil {
		g3.ZL().Fatal("AutoMigrate System Database", zap.Error(err))
	}
}

package system

import (
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/migrations"
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/crud"
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
		boot.Logger.Fatal("AutoMigrate System Database, err = %s", err.Error())
	}
}

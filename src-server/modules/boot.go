package modules

import (
	_ "github.com/zhouhp1295/g3-cms/modules/common/api"
	"github.com/zhouhp1295/g3-cms/modules/content"
	_ "github.com/zhouhp1295/g3-cms/modules/content/api"
	"github.com/zhouhp1295/g3-cms/modules/system"
	_ "github.com/zhouhp1295/g3-cms/modules/system/api"
)

func InitModulesDatabase() {
	//系统模块
	system.SyncTables()
	system.DoMigrate()
	//内容管理模块
	content.SyncTables()
	content.DoMigrate()
}

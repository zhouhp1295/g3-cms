package modules

import (
	_ "github.com/zhouhp1295/g3-cms/modules/common/api"
	"github.com/zhouhp1295/g3-cms/modules/system"
	_ "github.com/zhouhp1295/g3-cms/modules/system/api"
)

func InitModulesDatabase() {
	system.SyncTables()
	system.DoMigrate()
}

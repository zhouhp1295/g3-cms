// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package modules

import (
	"github.com/zhouhp1295/g3-cms/boot"
	_ "github.com/zhouhp1295/g3-cms/modules/common/routers"
	"github.com/zhouhp1295/g3-cms/modules/content"
	_ "github.com/zhouhp1295/g3-cms/modules/install/routers"
	_ "github.com/zhouhp1295/g3-cms/modules/render"
	"github.com/zhouhp1295/g3-cms/modules/system"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		//系统模块
		system.SyncTables()
		system.DoMigrate()
		//内容管理模块
		content.SyncTables()
		content.DoMigrate()
		//创建超级管理员
		dao.SysUserDao.CreateSuperUser("admin", "123456")
		//初始化权限
		dao.SysRoleDao.RefreshRolePerms()
	})
}

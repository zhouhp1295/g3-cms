// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package migrations

import (
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/modules/system/migrations"
	"github.com/zhouhp1295/g3/crud"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var contentMenuData20220531 = `
[
	{"id":2, "name":"内容管理", "title":"内容管理", "path":"/content", "type":"1", "icon": "nested", "component":"Layout", "sort":0},
	{"id":200, "pid":2, "name":"Article", "title":"文章管理", "path":"article", "type":"2", "icon": "documentation", "component":"content/article/index", "perms":"content:article:list", "sort":0},
	{"id":20001, "pid":200, "title":"文章查询", "type":"3", "perms":"content:article:query", "sort":0},
	{"id":20002, "pid":200, "title":"文章新增", "type":"3", "perms":"content:article:add", "sort":1},
	{"id":20003, "pid":200, "title":"文章编辑", "type":"3", "perms":"content:article:edit", "sort":2},
	{"id":20004, "pid":200, "title":"文章删除", "type":"3", "perms":"content:article:remove", "sort":3},
	{"id":201, "pid":2, "name":"Category", "title":"分类管理", "path":"category", "type":"2", "icon": "tree-table", "component":"content/category/index", "perms":"content:category:list", "sort":10},
	{"id":20101, "pid":201, "title":"分类查询", "type":"3", "perms":"content:category:query", "sort":0},
	{"id":20102, "pid":201, "title":"分类新增", "type":"3", "perms":"content:category:add", "sort":1},
	{"id":20103, "pid":201, "title":"分类编辑", "type":"3", "perms":"content:category:edit", "sort":2},
	{"id":20104, "pid":201, "title":"分类删除", "type":"3", "perms":"content:category:remove", "sort":3},
	{"id":202, "pid":2, "name":"Tag", "title":"标签管理", "path":"tag", "type":"2", "icon": "button", "component":"content/tag/index", "perms":"content:tag:list", "sort":20},
	{"id":20201, "pid":202, "title":"标签查询", "type":"3", "perms":"content:tag:query", "sort":0},
	{"id":20202, "pid":202, "title":"标签新增", "type":"3", "perms":"content:tag:add", "sort":1},
	{"id":20203, "pid":202, "title":"标签编辑", "type":"3", "perms":"content:tag:edit", "sort":2},
	{"id":20204, "pid":202, "title":"标签删除", "type":"3", "perms":"content:tag:remove", "sort":3},
	{"id":203, "pid":2, "name":"Menu", "title":"菜单管理", "path":"menu", "type":"2", "icon": "tree", "component":"content/menu/index", "perms":"content:menu:list", "sort":30},
	{"id":20301, "pid":203, "title":"菜单查询", "type":"3", "perms":"content:menu:query", "sort":0},
	{"id":20302, "pid":203, "title":"菜单新增", "type":"3", "perms":"content:menu:add", "sort":1},
	{"id":20303, "pid":203, "title":"菜单编辑", "type":"3", "perms":"content:menu:edit", "sort":2},
	{"id":20304, "pid":203, "title":"菜单删除", "type":"3", "perms":"content:menu:remove", "sort":3},
	{"id":204, "pid":2, "name":"Banner", "title":"Banner管理", "path":"banner", "type":"2", "icon": "guide", "component":"content/banner/index", "perms":"content:banner:list", "sort":40},
	{"id":20401, "pid":204, "title":"Banner查询", "type":"3", "perms":"content:banner:query", "sort":0},
	{"id":20402, "pid":204, "title":"Banner新增", "type":"3", "perms":"content:banner:add", "sort":1},
	{"id":20403, "pid":204, "title":"Banner编辑", "type":"3", "perms":"content:banner:edit", "sort":2},
	{"id":20404, "pid":204, "title":"Banner删除", "type":"3", "perms":"content:banner:remove", "sort":3},
	{"id":205, "pid":2, "name":"FriendLink", "title":"友链管理", "path":"friendLink", "type":"2", "icon": "link", "component":"content/friendLink/index", "perms":"content:friendLink:list", "sort":50},
	{"id":20501, "pid":205, "title":"友链查询", "type":"3", "perms":"content:friendLink:query", "sort":0},
	{"id":20502, "pid":205, "title":"友链新增", "type":"3", "perms":"content:friendLink:add", "sort":1},
	{"id":20503, "pid":205, "title":"友链编辑", "type":"3", "perms":"content:friendLink:edit", "sort":2},
	{"id":20504, "pid":205, "title":"友链删除", "type":"3", "perms":"content:friendLink:remove", "sort":3},
	{"id":206, "pid":2, "name":"ContentConfig", "title":"网站配置", "path":"config", "type":"2", "icon": "system", "component":"content/config/index", "perms":"content:config:list", "sort":60},
	{"id":20601, "pid":206, "title":"查看配置", "type":"3", "perms":"content:config:query", "sort":0},
	{"id":20602, "pid":206, "title":"编辑配置", "type":"3", "perms":"content:config:edit", "sort":1},
	{"id":207, "pid":2, "name":"Writer", "title":"作者管理", "path":"writer", "type":"2", "icon": "user", "component":"content/writer/index", "perms":"content:writer:list", "sort":9},
	{"id":20701, "pid":207, "title":"作者查询", "type":"3", "perms":"content:writer:query", "sort":0},
	{"id":20702, "pid":207, "title":"作者新增", "type":"3", "perms":"content:writer:add", "sort":1},
	{"id":20703, "pid":207, "title":"作者编辑", "type":"3", "perms":"content:writer:edit", "sort":2},
	{"id":20704, "pid":207, "title":"作者删除", "type":"3", "perms":"content:writer:remove", "sort":3}
]
`

const M20220531InitContentCode = "20220531_init_content"

func M20220531InitContent() func() error {
	return func() error {
		rootDB := crud.DbSess()
		//开启事务
		return rootDB.Transaction(func(tx *gorm.DB) error {
			err := migrations.CreateSystemMenus(tx, contentMenuData20220531)
			if err != nil {
				g3.ZL().Fatal("20220531_init_content", zap.Error(err))
				return err
			}
			return nil
		})
	}
}

package model

import "github.com/zhouhp1295/g3/crud"

type SysMenu struct {
	crud.BaseModel
	Pid       int64  `gorm:"NOT NULL;DEFAULT:0;COMMENT:'父节点ID'" json:"pid" form:"pid"`
	Name      string `gorm:"NOT NULL;TYPE:VARCHAR(20);COMMENT:'菜单名'" json:"name" form:"name"`
	Title     string `gorm:"NOT NULL;TYPE:VARCHAR(20);COMMENT:'菜单标题'" json:"title" form:"title"`
	Path      string `gorm:"TYPE:VARCHAR(255);COMMENT:'路由地址'" json:"path" form:"path"`
	Icon      string `gorm:"TYPE:VARCHAR(50);COMMENT:'图标'" json:"icon" form:"icon"`
	Sort      int    `gorm:"NOT NULL;DEFAULT:0;COMMENT:'排序'" json:"sort" form:"sort"`
	Component string `gorm:"TYPE:VARCHAR(50);COMMENT:'组件'" json:"component" form:"component"`
	Perms     string `gorm:"TYPE:VARCHAR(50);COMMENT:'权限'" json:"perms" form:"perms"`
	Query     string `gorm:"TYPE:VARCHAR(255);COMMENT:'路由参数'" json:"query" form:"query"`
	IsCache   string `gorm:"NOT NULL;TYPE:CHAR(1);DEFAULT:'1';COMMENT:'是否缓存:0=NO 1=YES'" json:"isCache" form:"isCache"`
	Visible   string `gorm:"NOT NULL;TYPE:CHAR(1);DEFAULT:'1';COMMENT:'是否显示:0=NO 1=YES'" json:"visible" form:"visible"  query:"eq"`
	IsFrame   string `gorm:"NOT NULL;TYPE:CHAR(1);DEFAULT:'0';COMMENT:'是否外联:0=NO 1=YES'" json:"isFrame" form:"isFrame"`
	crud.TailColumns
}

// Table 返回表名
func (*SysMenu) Table() string {
	return "sys_menu"
}

// NewModel 返回实例
func (*SysMenu) NewModel() crud.ModelInterface {
	return new(SysMenu)
}

// NewModels 返回实例数组
func (*SysMenu) NewModels() interface{} {
	return make([]SysMenu, 0)
}

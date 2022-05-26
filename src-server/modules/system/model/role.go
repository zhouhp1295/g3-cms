package model

import (
	"github.com/zhouhp1295/g3/crud"
)

type SysRole struct {
	crud.BaseModel
	Name       string `gorm:"NOT NULL;TYPE:VARCHAR(20);COMMENT:'角色名'" json:"name" form:"name"  query:"like"`
	Identifier string `gorm:"UNIQUE;TYPE:VARCHAR(20);COMMENT:'角色标识符'" json:"identifier" form:"identifier" query:"like"`
	Sort       int    `gorm:"NOT NULL;DEFAULT:0;COMMENT:'排序'" json:"sort" form:"sort"`
	Menus      string `gorm:"TYPE:TEXT;COMMENT:'权限菜单'" json:"menus" form:"menus"`
	crud.TailColumns
}

// Table 返回表名
func (*SysRole) Table() string {
	return "sys_role"
}

// NewModel 返回实例
func (*SysRole) NewModel() crud.ModelInterface {
	return new(SysRole)
}

// NewModels 返回实例数组
func (*SysRole) NewModels() interface{} {
	return make([]SysRole, 0)
}

// GetUpdateColumns 更新时的列
func (*SysRole) GetUpdateColumns() []string {
	return []string{"name", "identifier", "sort", "menus", "updated_by", "updated_at", "remark"}
}

package model

import (
	"github.com/zhouhp1295/g3/crud"
)

type SysConfig struct {
	crud.BaseModel
	Name  string `gorm:"TYPE:VARCHAR(100);COMMENT:'配置名'" json:"name" form:"name" query:"like"`
	Code  string `gorm:"TYPE:VARCHAR(100);COMMENT:'代码'"  json:"code" form:"code" query:"like"`
	Value string `gorm:"TYPE:VARCHAR(100);COMMENT:'值'"  json:"value" form:"value"`
	crud.TailColumns
}

// Table 返回表名
func (*SysConfig) Table() string {
	return "sys_config"
}

// NewModel 返回实例
func (*SysConfig) NewModel() crud.ModelInterface {
	return new(SysConfig)
}

// NewModels 返回实例数组
func (*SysConfig) NewModels() interface{} {
	return make([]SysConfig, 0)
}

// GetUpdateColumns 更新时的列
func (*SysConfig) GetUpdateColumns() []string {
	return []string{"name", "code", "value", "updated_by", "updated_at", "remark"}
}

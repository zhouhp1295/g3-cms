package model

import "github.com/zhouhp1295/g3/crud"

type SysDictType struct {
	crud.BaseModel
	Name string `gorm:"TYPE:VARCHAR(100);COMMENT:'字典名称'" json:"name" form:"name"`
	Code string `gorm:"TYPE:VARCHAR(100);COMMENT:'字典编号'" json:"code" form:"code"`
	crud.TailColumns
}

// Table 返回表名
func (*SysDictType) Table() string {
	return "sys_dict_type"
}

// NewModel 返回实例
func (*SysDictType) NewModel() crud.ModelInterface {
	return new(SysDictType)
}

// NewModels 返回实例数组
func (*SysDictType) NewModels() interface{} {
	return make([]SysDictType, 0)
}

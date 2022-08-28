// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package model

import "github.com/zhouhp1295/g3/crud"

type SysDictData struct {
	crud.BaseModel
	Label     string `gorm:"TYPE:VARCHAR(100);COMMENT:字典标签"  json:"label" form:"label"`
	Code      string `gorm:"TYPE:VARCHAR(100);COMMENT:字典类型" json:"code" form:"code"`
	Value     string `gorm:"TYPE:VARCHAR(100);COMMENT:字典键值" json:"value" form:"value"`
	IsDefault string `gorm:"TYPE:CHAR(1);NOT NULL;DEFAULT:0;COMMENT:状态 0=NO 1=YES" json:"isDefault"  form:"isDefault"`
	Sort      int    `gorm:"NOT NULL;DEFAULT:0;COMMENT:排序" json:"sort" form:"sort"`
	crud.TailColumns
}

// Table 返回表名
func (*SysDictData) Table() string {
	return "sys_dict_data"
}

// NewModel 返回实例
func (*SysDictData) NewModel() crud.ModelInterface {
	return new(SysDictData)
}

// NewModels 返回实例数组
func (*SysDictData) NewModels() interface{} {
	return make([]SysDictData, 0)
}

// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package model

import "github.com/zhouhp1295/g3/crud"

type ContentTag struct {
	crud.BaseModel
	Title       string `gorm:"TYPE:VARCHAR(100);INDEX;COMMENT:标题" json:"title" form:"title" query:"like"`
	Cover       string `gorm:"TYPE:VARCHAR(255);COMMENT:封面图片" json:"cover" form:"cover" query:"like"`
	Keywords    string `gorm:"TYPE:VARCHAR(100);COMMENT:关键字" json:"keywords" form:"keywords" query:"like"`
	Description string `gorm:"TYPE:VARCHAR(100);COMMENT:描述" json:"description" form:"description" query:"like"`
	Sort        int    `gorm:"NOT NULL;DEFAULT:0;COMMENT:排序" json:"sort" form:"sort"`
	crud.TailColumns
}

// Table 返回表名
func (*ContentTag) Table() string {
	return "content_tag"
}

// NewModel 返回实例
func (*ContentTag) NewModel() crud.ModelInterface {
	return new(ContentTag)
}

// NewModels 返回实例数组
func (*ContentTag) NewModels() interface{} {
	return make([]ContentTag, 0)
}

// GetUpdateColumns 更新时的列
func (*ContentTag) GetUpdateColumns() []string {
	return []string{"title", "cover", "keywords", "description", "sort", "updated_by", "updated_at", "remark"}
}

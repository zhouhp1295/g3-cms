// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package model

import "github.com/zhouhp1295/g3/crud"

type ContentBanner struct {
	crud.BaseModel
	Title string `gorm:"TYPE:VARCHAR(100);COMMENT:标题" json:"title" form:"title" query:"like"`
	Cover string `gorm:"TYPE:VARCHAR(100);COMMENT:图片" json:"cover" form:"cover"`
	Url   string `gorm:"TYPE:VARCHAR(100);COMMENT:值" json:"url" form:"url" query:"like"`
	Sort  int    `gorm:"NOT NULL;DEFAULT:0;COMMENT:排序" json:"sort" form:"sort"`
	crud.TailColumns
}

// Table 返回表名
func (*ContentBanner) Table() string {
	return "content_banner"
}

// NewModel 返回实例
func (*ContentBanner) NewModel() crud.ModelInterface {
	return new(ContentBanner)
}

// NewModels 返回实例数组
func (*ContentBanner) NewModels() interface{} {
	return make([]ContentBanner, 0)
}

// GetUpdateColumns 更新时的列
func (*ContentBanner) GetUpdateColumns() []string {
	return []string{"title", "cover", "url",
		"sort", "status", "updated_by", "updated_at", "remark"}
}

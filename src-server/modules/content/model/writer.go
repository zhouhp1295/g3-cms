// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved model

package model

import "github.com/zhouhp1295/g3/crud"

type ContentWriter struct {
	crud.BaseModel
	Name        string `gorm:"TYPE:VARCHAR(100);INDEX;COMMENT:作者名" json:"name" form:"name" query:"like"`
	Avatar      string `gorm:"TYPE:VARCHAR(255);COMMENT:头像图片" json:"avatar" form:"avatar"`
	Cover       string `gorm:"TYPE:VARCHAR(255);COMMENT:封面图片" json:"cover" form:"cover"`
	Sex         string `gorm:"TYPE:CHAR(1);NOT NULL;DEFAULT:0;COMMENT:性别" json:"sex" form:"sex" query:"eq"`
	Keywords    string `gorm:"TYPE:VARCHAR(100);COMMENT:关键字" json:"keywords" form:"keywords"`
	Description string `gorm:"TYPE:VARCHAR(100);COMMENT:描述" json:"description" form:"description"`
	NumArticle  int64  `gorm:"NOT NULL;DEFAULT:0;COMMENT:文章数量" json:"numArticle" form:"numArticle"`
	Sort        int    `gorm:"NOT NULL;DEFAULT:0;COMMENT:排序" json:"sort" form:"sort"`
	crud.TailColumns
}

// Table 返回表名
func (*ContentWriter) Table() string {
	return "content_writer"
}

// NewModel 返回实例
func (*ContentWriter) NewModel() crud.ModelInterface {
	return new(ContentWriter)
}

// NewModels 返回实例数组
func (*ContentWriter) NewModels() interface{} {
	return make([]ContentWriter, 0)
}

// GetUpdateColumns 更新时的列
func (*ContentWriter) GetUpdateColumns() []string {
	return []string{"name", "avatar", "cover", "sex", "keywords", "description", "num_article", "sort", "updated_by", "updated_at", "remark"}
}

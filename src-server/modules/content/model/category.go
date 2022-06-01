package model

import "github.com/zhouhp1295/g3/crud"

type ContentCategory struct {
	crud.BaseModel
	Pid          int64  `gorm:"NOT NULL;DEFAULT:0;COMMENT:父ID" json:"pid" form:"pid" query:"eq"`
	Ancestors    string `gorm:"TYPE:VARCHAR(100);DEFAULT:0,;COMMENT:祖级列表" json:"ancestors" form:"ancestors" query:"like"`
	Title        string `gorm:"TYPE:VARCHAR(100);INDEX;COMMENT:标题" json:"title" form:"title" query:"like"`
	Cover        string `gorm:"TYPE:VARCHAR(255);COMMENT:封面图片" json:"cover" form:"cover"`
	Icon         string `gorm:"TYPE:VARCHAR(255);COMMENT:Icon" json:"icon" form:"icon"`
	Banner       string `gorm:"TYPE:VARCHAR(255);COMMENT:Banner图片" json:"banner" form:"banner"`
	Keywords     string `gorm:"TYPE:VARCHAR(100);COMMENT:关键字" json:"keywords" form:"keywords" query:"like"`
	Description  string `gorm:"TYPE:VARCHAR(100);COMMENT:描述" json:"description" form:"description" query:"like"`
	Sort         int    `gorm:"NOT NULL;DEFAULT:0;COMMENT:排序" json:"sort" form:"sort"`
	InMenu       string `gorm:"TYPE:CHAR(1);NOT NULL;DEFAULT:0;COMMENT:显示到菜单 0=NO 1=YES" json:"inMenu" form:"inMenu"  query:"eq"`
	InMenuSort   int    `gorm:"NOT NULL;DEFAULT:0;COMMENT:排序" json:"inMenuSort" form:"inMenuSort"`
	InBanner     string `gorm:"TYPE:CHAR(1);NOT NULL;DEFAULT:0;COMMENT:显示在菜单 0=NO 1=YES" json:"inBanner" form:"inBanner" query:"eq"`
	InBannerSort int    `gorm:"NOT NULL;DEFAULT:0;COMMENT:排序" json:"inBannerSort" form:"inBannerSort"`
	crud.TailColumns
}

// Table 返回表名
func (*ContentCategory) Table() string {
	return "content_category"
}

// NewModel 返回实例
func (*ContentCategory) NewModel() crud.ModelInterface {
	return new(ContentCategory)
}

// NewModels 返回实例数组
func (*ContentCategory) NewModels() interface{} {
	return make([]ContentCategory, 0)
}

// GetUpdateColumns 更新时的列
func (*ContentCategory) GetUpdateColumns() []string {
	return []string{"pid", "ancestors", "title", "cover", "icon", "banner",
		"keywords", "description", "sort", "in_menu", "in_menu_sort", "in_banner", "in_banner_sort",
		"status", "updated_by", "updated_at", "remark"}
}

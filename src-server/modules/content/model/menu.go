package model

import "github.com/zhouhp1295/g3/crud"

type ContentMenu struct {
	crud.BaseModel
	Pid       int64  `gorm:"NOT NULL;DEFAULT:0;COMMENT:父节点ID" json:"pid" form:"pid" query:"eq"`
	Ancestors string `gorm:"TYPE:VARCHAR(100);DEFAULT:0,;COMMENT:祖级列表" json:"ancestors" form:"ancestors" query:"like"`
	Title     string `gorm:"NOT NULL;TYPE:VARCHAR(20);COMMENT:菜单标题" json:"title" form:"title" query:"like"`
	Url       string `gorm:"TYPE:VARCHAR(255);COMMENT:URL地址" json:"url" form:"url"`
	Icon      string `gorm:"TYPE:VARCHAR(255);COMMENT:Icon" json:"icon" form:"icon"`
	Sort      int    `gorm:"NOT NULL;DEFAULT:0;COMMENT:排序" json:"sort" form:"sort"`
	IsBlank   string `gorm:"NOT NULL;TYPE:CHAR(1);DEFAULT:1;COMMENT:是否新页面 0=NO 1=YES" json:"isBlank" form:"isBlank"  query:"eq"`
	crud.TailColumns
}

// Table 返回表名
func (*ContentMenu) Table() string {
	return "content_menu"
}

// NewModel 返回实例
func (*ContentMenu) NewModel() crud.ModelInterface {
	return new(ContentMenu)
}

// NewModels 返回实例数组
func (*ContentMenu) NewModels() interface{} {
	return make([]ContentMenu, 0)
}

// GetUpdateColumns 更新时的列
func (*ContentMenu) GetUpdateColumns() []string {
	return []string{"pid", "ancestors", "title", "url", "icon", "is_blank", "sort", "status", "updated_by", "updated_at", "remark"}
}

package model

import "github.com/zhouhp1295/g3/crud"

type ContentArticleTag struct {
	crud.BaseModel
	ArticleId int64 `gorm:"NOT NULL;INDEX;COMMENT:文章ID" json:"articleId" json:"articleId"`
	TagId     int64 `gorm:"NOT NULL;INDEX;COMMENT:标签ID" json:"tagId" json:"tagId"`
	Cnt       int64 `gorm:"-"`
	crud.TailColumns
}

// Table 返回表名
func (*ContentArticleTag) Table() string {
	return "content_article_tag"
}

// NewModel 返回实例
func (*ContentArticleTag) NewModel() crud.ModelInterface {
	return new(ContentArticleTag)
}

// NewModels 返回实例数组
func (*ContentArticleTag) NewModels() interface{} {
	return make([]ContentArticleTag, 0)
}

// GetUpdateColumns 更新时的列
func (*ContentArticleTag) GetUpdateColumns() []string {
	return []string{"article_id", "tag_id", "status", "updated_by", "updated_at", "remark"}
}

package model

import (
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
	"time"
)

type ContentArticle struct {
	crud.BaseModel
	Title           string                 `gorm:"TYPE:VARCHAR(100);INDEX;COMMENT:标题" json:"title" form:"title" query:"like"`
	Category        int64                  `gorm:"NOT NULL;INDEX;DEFAULT:0;COMMENT:分类" json:"category" form:"category" query:"eq"`
	Cover           string                 `gorm:"TYPE:VARCHAR(255);COMMENT:封面图片" json:"cover" form:"cover"`
	Banner          string                 `gorm:"TYPE:VARCHAR(255);COMMENT:Banner图片" json:"banner" form:"banner"`
	SeoTitle        string                 `gorm:"TYPE:VARCHAR(100);INDEX;COMMENT:标题" json:"seoTitle" form:"seoTitle" query:"like"`
	SeoKeywords     string                 `gorm:"TYPE:VARCHAR(255);COMMENT:关键字" json:"seoKeywords" form:"seoKeywords" query:"like"`
	SeoDescription  string                 `gorm:"TYPE:VARCHAR(255);COMMENT:描述" json:"seoDescription" form:"seoDescription" query:"like"`
	Content         string                 `gorm:"TYPE:TEXT;COMMENT:内容" json:"content" form:"content"`
	NumRead         int64                  `gorm:"NOT NULL;DEFAULT:0;COMMENT:阅读量" json:"numRead" form:"numRead"`
	IsOriginal      string                 `gorm:"TYPE:CHAR(1);NOT NULL;DEFAULT:0;COMMENT:原创 0=NO 1=YES" json:"isOriginal" form:"isOriginal" query:"eq"`
	InBanner        string                 `gorm:"TYPE:CHAR(1);NOT NULL;DEFAULT:0;COMMENT:显示到Banner 0=NO 1=YES" json:"inBanner" form:"inBanner" query:"eq"`
	InBannerSort    int                    `gorm:"NOT NULL;DEFAULT:0;COMMENT:排序" json:"inBannerSort" form:"inBannerSort" query:"eq"`
	Sort            int                    `gorm:"NOT NULL;DEFAULT:0;COMMENT:排序" json:"sort" form:"sort"`
	Writer          int64                  `gorm:"NOT NULL;DEFAULT:0;COMMENT:作者" json:"writer" form:"writer" query:"eq"`
	PublishedBy     int64                  `gorm:"NOT NULL;DEFAULT:0;COMMENT:发布人" json:"publishedBy" form:"publishedBy" `
	PublishedAt     time.Time              `gorm:"COMMENT:发布时间" json:"publishedAt" form:"publishedAt"`
	Tags            []int64                `gorm:"-" json:"tags" form:"tags"`
	TagOptions      []helpers.SelectOption `gorm:"-" json:"tagOptions" form:"tagOptions"`
	WriterName      string                 `gorm:"-"`
	FrontCategories []helpers.TreeOption   `gorm:"-"`
	CategoryName    string                 `gorm:"-"`
	crud.TailColumns
}

// Table 返回表名
func (*ContentArticle) Table() string {
	return "content_article"
}

// NewModel 返回实例
func (*ContentArticle) NewModel() crud.ModelInterface {
	return new(ContentArticle)
}

// NewModels 返回实例数组
func (*ContentArticle) NewModels() interface{} {
	return make([]ContentArticle, 0)
}

// GetUpdateColumns 更新时的列
func (*ContentArticle) GetUpdateColumns() []string {
	return []string{"title", "category", "cover", "banner", "seo_title", "seo_keywords", "seo_description",
		"content", "num_read", "in_banner", "in_banner_sort", "writer", "published_by", "published_at",
		"sort", "status", "updated_by", "updated_at", "remark"}
}

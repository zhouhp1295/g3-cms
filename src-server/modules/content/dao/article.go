package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
)

type contentArticleDAO struct {
	crud.BaseDao
}

var ContentArticleDao = &contentArticleDAO{
	crud.BaseDao{Model: new(model.ContentArticle)},
}

type FrontArticleSimpleData struct {
	Id           int64                  `json:"id"`
	Title        string                 `json:"title"`
	SeoTitle     string                 `json:"seoTitle"`
	Cover        string                 `json:"cover"`
	Excerpt      string                 `json:"excerpt"`
	Writer       int64                  `json:"writer"`
	WriterName   string                 `json:"writerName"`
	NumRead      int64                  `json:"numRead"`
	PublishedAt  string                 `json:"publishedAt"`
	Category     int64                  `json:"category"`
	CategoryName string                 `json:"categoryName"`
	Tags         []helpers.SelectOption `json:"tags"`
}

type FrontArticleDetailData struct {
	FrontArticleSimpleData
	SeoKeywords    string `json:"seoKeywords"`
	SeoDescription string `json:"seoDescription"`
	Content        string `json:"content"`
}

func (dao *contentArticleDAO) AfterGet(m crud.ModelInterface) {
	if _m, _ok := m.(*model.ContentArticle); _ok {
		tagOptions := make([]helpers.SelectOption, 0)
		query := crud.DbSess().Table("content_article_tag").
			Joins("left join content_tag on content_article_tag.tag_id = content_tag.id").
			Select([]string{"content_tag.id", "content_tag.title as label"}).
			Where("content_article_tag.article_id = ? and content_article_tag.deleted = ? and content_tag.deleted = ?", _m.Id, crud.FlagNo, crud.FlagNo)
		query.Find(&tagOptions)

		if len(tagOptions) > 0 {
			for _, tag := range tagOptions {
				_m.Tags = append(_m.Tags, tag.Id)
			}
		} else {
			_m.Tags = make([]int64, 0)
		}

		_m.TagOptions = tagOptions
	}
}

//
//func (dao *contentArticleDAO) BeforeInsert(m crud.ModelInterface) (ok bool, msg string) {
//	return
//}

func (dao *contentArticleDAO) AfterInsert(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.ContentArticle); _ok && len(_m.Tags) > 0 {
		tags := make([]model.ContentArticleTag, len(_m.Tags))
		for i, tagId := range _m.Tags {
			tags[i] = model.ContentArticleTag{
				ArticleId: _m.Id,
				TagId:     tagId,
			}
		}
		crud.DbSess().CreateInBatches(tags, len(tags))
	}
	ok = true
	return
}

//
//func (dao *contentArticleDAO) BeforeUpdate(m crud.ModelInterface) (ok bool, msg string) {
//	return
//}

func (dao *contentArticleDAO) AfterUpdate(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.ContentArticle); _ok && len(_m.Tags) > 0 {

		existTags := make([]model.ContentArticleTag, 0)

		crud.DbSess().Where("article_id = ? and deleted =?", _m.Id, crud.FlagNo).Find(&existTags)

		for _, existTag := range existTags {
			idx := helpers.IndexOf[int64](_m.Tags, existTag.TagId)
			if idx >= 0 {
				_m.Tags[idx] = 0
				continue
			}
			crud.DbSess().Delete(existTag)
		}

		insertTags := make([]model.ContentArticleTag, 0)
		for _, tagId := range _m.Tags {
			if tagId > 0 {
				insertTags = append(insertTags, model.ContentArticleTag{
					ArticleId: _m.Id,
					TagId:     tagId,
				})
			}
		}
		if len(insertTags) > 0 {
			crud.DbSess().CreateInBatches(insertTags, len(insertTags))
		}
	}
	ok = true
	return
}

//
//func (dao *contentArticleDAO) BeforeDelete(m crud.ModelInterface) (ok bool, msg string) {
//	return
//}
//
//func (dao *contentArticleDAO) BeforeRemove(m crud.ModelInterface) (ok bool, msg string) {
//	return
//}

func (dao *contentArticleDAO) FrontTopArticles() []FrontArticleSimpleData {
	return getFrontTopArticlesFromCache()
}

func (dao *contentArticleDAO) FrontLatestArticles(page int) ([]FrontArticleSimpleData, crud.PageData) {
	pageSize := 20

	if page < 1 {
		page = 1
	}
	var total int64
	crud.DbSess().Where("status = ? and deleted  = ?", crud.FlagYes, crud.FlagNo).Table("content_article").Count(&total)
	if total == 0 {
		return make([]FrontArticleSimpleData, 0), crud.PageResult(page, pageSize, int(total))
	}

	articleRows := make([]model.ContentArticle, 0)
	crud.DbSess().Where("status = ? and deleted  = ?", crud.FlagYes, crud.FlagNo).
		Order("published_at desc").
		Limit(pageSize).Offset((page - 1) * pageSize).
		Find(&articleRows)

	articles := make([]FrontArticleSimpleData, len(articleRows))

	for i, article := range articleRows {
		_data := FrontArticleSimpleData{
			Id:       article.Id,
			Title:    article.Title,
			SeoTitle: article.SeoTitle,
			Cover:    article.Cover,
			Writer:   article.Writer,
			NumRead:  article.NumRead,
			Category: article.Category,
		}

		if len(_data.SeoTitle) == 0 {
			_data.SeoTitle = _data.Title
		}

		_data.WriterName = getWriterName(_data.Writer)

		_data.CategoryName = getCategoryName(_data.Category)

		articles[i] = _data
	}

	return articles, crud.PageResult(page, pageSize, int(total))
}

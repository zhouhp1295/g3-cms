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

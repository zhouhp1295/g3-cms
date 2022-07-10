package dao

import (
	"fmt"
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
	"gorm.io/gorm"
	"time"
)

type contentArticleDAO struct {
	crud.BaseDao
}

var ContentArticleDao = &contentArticleDAO{
	crud.BaseDao{Model: new(model.ContentArticle)},
}

func (dao *contentArticleDAO) ClearCache() {
	clearFrontAllBannerCache()
	clearFrontAllArticleCache()
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

func (dao *contentArticleDAO) BeforeInsert(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.ContentArticle); _ok && _m.Status == crud.FlagYes {
		_m.PublishedBy = _m.CreatedBy
		_m.PublishedAt = time.Now()
	}
	ok = true
	return
}

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
	dao.ClearCache()
	return
}

func (dao *contentArticleDAO) BeforeUpdate(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.ContentArticle); _ok && _m.Status == crud.FlagYes {
		lastM := dao.FindByPk(_m.Id)
		_lastM, _ := lastM.(*model.ContentArticle)
		if _lastM.Status != _m.Status {
			_m.PublishedBy = _m.CreatedBy
			_m.PublishedAt = time.Now()
		}
	}
	ok = true
	return
}

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
	dao.ClearCache()
	return
}

func (dao *contentArticleDAO) UpdateStatus(pk int64, status interface{}, operator int64) bool {
	m := dao.FindByPk(pk)
	_m, _ := m.(*model.ContentArticle)
	_status, _ := status.(string)
	if _m.Status != _status && _status == crud.FlagYes {
		_m.PublishedBy = operator
		_m.PublishedAt = time.Now()
	}
	_m.Status = _status
	dao.Update(_m, operator)
	dao.ClearCache()
	return true
}

func (dao *contentArticleDAO) UpdateInBanner(pk int64, inBanner string, inBannerSort int, operator int64) bool {
	article := new(model.ContentArticle)
	article.Id = pk
	article.InBanner = inBanner
	article.InBannerSort = inBannerSort
	article.UpdatedBy = operator

	err := crud.DbSess().Select([]string{"in_banner", "in_banner_sort", "updated_by", "updated_at"}).Updates(article).Error

	if err != nil {
		return false
	}
	dao.ClearCache()
	return true
}

//
//func (dao *contentArticleDAO) BeforeDelete(m crud.ModelInterface) (ok bool, msg string) {
//	return
//}
//
//func (dao *contentArticleDAO) BeforeRemove(m crud.ModelInterface) (ok bool, msg string) {
//	return
//}

func (dao *contentArticleDAO) ApplyFrontData(m *model.ContentArticle) {
	applyFrontData(m)
}

func (dao *contentArticleDAO) FrontTopArticles() []model.ContentArticle {
	return getFrontTopArticlesFromCache()
}

func (dao *contentArticleDAO) FrontLatestArticles(page int) ([]model.ContentArticle, crud.PageData) {
	pageSize := crud.DefaultPageSize

	if page < 1 {
		page = 1
	}
	var total int64
	crud.DbSess().Where("status = ? and deleted  = ?", crud.FlagYes, crud.FlagNo).Table("content_article").Count(&total)
	if total == 0 {
		return make([]model.ContentArticle, 0), crud.PageResult(page, pageSize, int(total))
	}

	articleRows := make([]model.ContentArticle, 0)
	crud.DbSess().Where("status = ? and deleted  = ?", crud.FlagYes, crud.FlagNo).
		Order("published_at desc").
		Limit(pageSize).Offset((page - 1) * pageSize).
		Find(&articleRows)

	for _, article := range articleRows {

		if len(article.SeoTitle) == 0 {
			article.SeoTitle = article.Title
		}
		applyFrontData(&article)
	}

	return articleRows, crud.PageResult(page, pageSize, int(total))
}

func (dao *contentArticleDAO) FrontRightLatestHotArticles(categoryId int64) []model.ContentArticle {
	articleRows := make([]model.ContentArticle, 0)
	query := crud.DbSess().Where("status = ? and deleted  = ?",
		crud.FlagYes, crud.FlagNo)

	if categoryId > 0 {
		query.Order(fmt.Sprintf("abs(category - %d) asc,num_read desc,sort desc,published_at desc", categoryId))
	} else {
		query.Order("num_read desc,sort desc,published_at desc")
	}

	query.Limit(10).Find(&articleRows)

	for _, article := range articleRows {

		if len(article.SeoTitle) == 0 {
			article.SeoTitle = article.Title
		}
		applyFrontData(&article)
	}
	return articleRows
}

func (dao *contentArticleDAO) FrontRightRecommendArticles(categoryId int64) []model.ContentArticle {
	articleRows := make([]model.ContentArticle, 0)
	query := crud.DbSess().Where("status = ? and deleted  = ?",
		crud.FlagYes, crud.FlagNo)

	if categoryId > 0 {
		query.Order(fmt.Sprintf("abs(category - %d) asc,sort desc,num_read desc,published_at desc", categoryId))
	} else {
		query.Order("sort desc,num_read desc,published_at desc")
	}

	query.Limit(10).Find(&articleRows)

	for _, article := range articleRows {

		if len(article.SeoTitle) == 0 {
			article.SeoTitle = article.Title
		}
		applyFrontData(&article)
	}
	return articleRows
}

func (dao *contentArticleDAO) FrontTagRightLatestHotArticles(tagId int64) []model.ContentArticle {
	articleRows := make([]model.ContentArticle, 0)
	query := crud.DbSess().
		Model(new(model.ContentArticle)).
		Select("distinct content_article.*").
		Where("content_article.status = ? and content_article.deleted  = ?",
			crud.FlagYes, crud.FlagNo)

	if tagId > 0 {
		query.Joins("left join content_article_tag on content_article_tag.article_id = content_article.id").
			Order(fmt.Sprintf("abs(content_article_tag.tag_id - %d) asc,content_article.num_read desc,content_article.sort desc,content_article.published_at desc", tagId))
	} else {
		query.Order("content_article.num_read desc,content_article.sort desc,content_article.published_at desc")
	}

	query.Limit(10).Find(&articleRows)

	for _, article := range articleRows {

		if len(article.SeoTitle) == 0 {
			article.SeoTitle = article.Title
		}
		applyFrontData(&article)
	}
	return articleRows
}

func (dao *contentArticleDAO) FrontTagRightRecommendArticles(tagId int64) []model.ContentArticle {
	articleRows := make([]model.ContentArticle, 0)
	query := crud.DbSess().
		Model(new(model.ContentArticle)).
		Select("distinct content_article.*").
		Where("content_article.status = ? and content_article.deleted  = ?",
			crud.FlagYes, crud.FlagNo)
	if tagId > 0 {
		query.Joins("left join content_article_tag on content_article_tag.article_id = content_article.id").
			Order(fmt.Sprintf("abs(content_article_tag.tag_id - %d) asc,content_article.num_read desc,content_article.sort desc,content_article.published_at desc", tagId))
	} else {
		query.Order("content_article.num_read desc,content_article.sort desc,content_article.published_at desc")
	}

	query.Limit(10).Find(&articleRows)

	for _, article := range articleRows {

		if len(article.SeoTitle) == 0 {
			article.SeoTitle = article.Title
		}
		applyFrontData(&article)
	}
	return articleRows
}

func (dao *contentArticleDAO) FrontWriterRightLatestHotArticles(writerId int64) []model.ContentArticle {
	articleRows := make([]model.ContentArticle, 0)
	query := crud.DbSess().Where("status = ? and deleted  = ?",
		crud.FlagYes, crud.FlagNo)

	if writerId > 0 {
		query.Order(fmt.Sprintf("abs(writer - %d) asc,num_read desc,sort desc,published_at desc", writerId))
	} else {
		query.Order("num_read desc,sort desc,published_at desc")
	}

	query.Limit(10).Find(&articleRows)

	for _, article := range articleRows {

		if len(article.SeoTitle) == 0 {
			article.SeoTitle = article.Title
		}
		applyFrontData(&article)
	}
	return articleRows
}

func (dao *contentArticleDAO) FrontWriterRightRecommendArticles(writerId int64) []model.ContentArticle {
	articleRows := make([]model.ContentArticle, 0)
	query := crud.DbSess().Where("status = ? and deleted  = ?",
		crud.FlagYes, crud.FlagNo)

	if writerId > 0 {
		query.Order(fmt.Sprintf("abs(writer - %d) asc,sort desc,num_read desc,published_at desc", writerId))
	} else {
		query.Order("sort desc,num_read desc,published_at desc")
	}

	query.Limit(10).Find(&articleRows)

	for _, article := range articleRows {

		if len(article.SeoTitle) == 0 {
			article.SeoTitle = article.Title
		}
		applyFrontData(&article)
	}
	return articleRows
}

func (dao *contentArticleDAO) FrontCategoryArticles(categoryId int64, page int) ([]model.ContentArticle, crud.PageData) {
	pageSize := crud.DefaultPageSize

	if page < 1 {
		page = 1
	}
	treeOptions := listFrontCategoryTreeOptionsFromCache()
	categoryOptions := helpers.FindChildrenIdList(treeOptions, categoryId)
	if categoryId > 0 {
		categoryOptions = append(categoryOptions, categoryId)
	}

	findScope := func(db *gorm.DB) *gorm.DB {
		db.Where("status = ? and deleted  = ?", crud.FlagYes, crud.FlagNo)
		if len(categoryOptions) > 0 {
			db.Where("category in (?)", categoryOptions)
		}
		return db
	}

	var total int64
	crud.DbSess().Scopes(findScope).Table("content_article").Count(&total)
	if total == 0 {
		return make([]model.ContentArticle, 0), crud.PageResult(page, pageSize, int(total))
	}

	articleRows := make([]model.ContentArticle, 0)
	crud.DbSess().Scopes(findScope).
		Order("published_at desc").
		Limit(pageSize).Offset((page - 1) * pageSize).
		Find(&articleRows)

	for _, article := range articleRows {
		if len(article.SeoTitle) == 0 {
			article.SeoTitle = article.Title
		}
		applyFrontData(&article)
	}

	return articleRows, crud.PageResult(page, pageSize, int(total))
}

func (dao *contentArticleDAO) FrontWriterArticles(writerId int64, page int) ([]model.ContentArticle, crud.PageData) {
	pageSize := crud.DefaultPageSize

	if page < 1 {
		page = 1
	}

	findScope := func(db *gorm.DB) *gorm.DB {
		db.Where("status = ? and deleted  = ?", crud.FlagYes, crud.FlagNo)
		if writerId > 0 {
			db.Where("writer = ?", writerId)
		}
		return db
	}

	var total int64
	crud.DbSess().Scopes(findScope).Table("content_article").Count(&total)
	if total == 0 {
		return make([]model.ContentArticle, 0), crud.PageResult(page, pageSize, int(total))
	}

	articleRows := make([]model.ContentArticle, 0)
	crud.DbSess().Scopes(findScope).
		Order("published_at desc").
		Limit(pageSize).Offset((page - 1) * pageSize).
		Find(&articleRows)

	for _, article := range articleRows {
		if len(article.SeoTitle) == 0 {
			article.SeoTitle = article.Title
		}
		applyFrontData(&article)
	}

	return articleRows, crud.PageResult(page, pageSize, int(total))
}

func (dao *contentArticleDAO) FrontTagArticles(tagId int64, page int) ([]model.ContentArticle, crud.PageData) {
	pageSize := crud.DefaultPageSize

	if page < 1 {
		page = 1
	}

	findScope := func(db *gorm.DB) *gorm.DB {
		db.Where("content_article.status = ? and content_article.deleted  = ?", crud.FlagYes, crud.FlagNo)
		if tagId > 0 {
			db.Joins("left join content_article_tag on content_article_tag.article_id = content_article.id")
			db.Where("content_article_tag.tag_id = ?", tagId)
		}
		return db
	}

	var total int64
	crud.DbSess().Model(new(model.ContentArticle)).Scopes(findScope).Count(&total)
	if total == 0 {
		return make([]model.ContentArticle, 0), crud.PageResult(page, pageSize, int(total))
	}

	articleRows := make([]model.ContentArticle, 0)
	crud.DbSess().Model(new(model.ContentArticle)).Select("content_article.*").Scopes(findScope).
		Order("content_article.published_at desc").
		Limit(pageSize).Offset((page - 1) * pageSize).
		Find(&articleRows)

	for _, article := range articleRows {
		if len(article.SeoTitle) == 0 {
			article.SeoTitle = article.Title
		}
		applyFrontData(&article)
	}

	return articleRows, crud.PageResult(page, pageSize, int(total))
}

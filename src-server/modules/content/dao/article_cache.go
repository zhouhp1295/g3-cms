// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package dao

import (
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
	"time"
)

func clearFrontAllArticleCache() {
	boot.Lache.Delete("K-Content-Dao-Article-FrontTopArticles")
}

func applyFrontData(m *model.ContentArticle) {
	m.WriterName = getWriterName(m.Writer)
	m.CategoryName = getCategoryName(m.Category)
	m.FrontCategories = ContentCategoryDao.GetFrontCategory(m.Category)
}

// 根据一定的逻辑取固定个数的最热文章
func getFrontTopArticles() (articles []model.ContentArticle, ok bool) {
	//逻辑: 最近一周点击量最多的一篇三篇文章
	limit := 3
	articles = make([]model.ContentArticle, 0)
	crud.DbSess().Where("status = ? and deleted  = ?", crud.FlagYes, crud.FlagNo).
		Order("num_read desc,published_at desc").
		Limit(limit).
		Find(&articles)

	if len(articles) == 0 {
		ok = false
		return
	}

	for _, article := range articles {
		if len(article.SeoTitle) == 0 {
			article.SeoTitle = article.Title
		}
		applyFrontData(&article)
	}
	ok = true
	return
}

// 根据一定的逻辑取固定个数的最热文章
func getFrontTopArticlesFromCache() []model.ContentArticle {
	key := "K-Content-Dao-Article-FrontTopArticles"
	result := make([]model.ContentArticle, 0)
	ok := boot.Lache.GetT(key, &result)
	if !ok {
		_data, _ok := getFrontTopArticles()
		if _ok {
			boot.Lache.Set(key, _data, 30*time.Minute)
		} else {
			boot.Lache.Set(key, _data, 10*time.Second)
		}
		return _data
	}
	return result
}

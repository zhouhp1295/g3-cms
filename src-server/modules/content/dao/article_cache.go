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

// 根据一定的逻辑取固定个数的最热文章
func getFrontTopArticles() (articles []FrontArticleSimpleData, ok bool) {
	//逻辑: 最近一周点击量最多的一篇三篇文章
	limit := 3
	articleRows := make([]model.ContentArticle, 0)
	crud.DbSess().Where("status = ? and deleted  = ?", crud.FlagYes, crud.FlagNo).
		Order("num_read desc,published_at desc").
		Limit(limit).Offset(0).
		Find(&articleRows)

	articles = make([]FrontArticleSimpleData, len(articleRows))

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

	return
}

// 根据一定的逻辑取固定个数的最热文章
func getFrontTopArticlesFromCache() []FrontArticleSimpleData {
	key := "K-Content-Dao-Article-FrontTopArticles"
	result := make([]FrontArticleSimpleData, 0)
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

package page

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
)

func CategoryHandler(ctx *gin.Context) {
	view, err := jetEngine().Set.GetTemplate("category.html")
	if err != nil {
		boot.Logger.Error("CategoryHandler, err = %s", err.Error())
		errorPage(ctx, err.Error())
		return
	}
	page := parseIntParam(ctx, "page", 1)
	categoryId := int64(parseIntParam(ctx, "categoryId", 0))
	category := dao.ContentCategoryDao.FindByPk(categoryId)
	articles, pageData := dao.ContentArticleDao.FrontCategoryArticles(categoryId, page)

	data := globalVarMap()
	data.Set("category", category)
	data.Set("articles", articles)
	data.Set("pageUrl", "category")
	data.Set("pageData", pageData)
	data.Set("rightLatestHotArticles", dao.ContentArticleDao.FrontRightLatestHotArticles(categoryId))
	data.Set("rightRecommendArticles", dao.ContentArticleDao.FrontRightRecommendArticles(categoryId))

	err = view.Execute(ctx.Writer, data, nil)

	if err != nil {
		boot.Logger.Error("CategoryHandler View Execute, err =", err.Error())
		errorPage(ctx, err.Error())
	}
}

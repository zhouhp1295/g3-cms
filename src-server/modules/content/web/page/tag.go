package page

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
)

func TagHandler(ctx *gin.Context) {
	view, err := jetEngine().Set.GetTemplate("tag.html")
	if err != nil {
		boot.Logger.Error("TagHandler, err = %s", err.Error())
		errorPage(ctx, err.Error())
		return
	}
	page := parseIntParam(ctx, "page", 1)
	tagId := int64(parseIntParam(ctx, "tagId", 0))
	tag := dao.ContentCategoryDao.FindByPk(tagId)
	articles, pageData := dao.ContentArticleDao.FrontTagArticles(tagId, page)

	data := globalVarMap()
	data.Set("tag", tag)
	data.Set("articles", articles)
	data.Set("pageUrl", "category")
	data.Set("pageData", pageData)
	data.Set("rightLatestHotArticles", dao.ContentArticleDao.FrontTagRightLatestHotArticles(tagId))
	data.Set("rightRecommendArticles", dao.ContentArticleDao.FrontTagRightRecommendArticles(tagId))

	err = view.Execute(ctx.Writer, data, nil)

	if err != nil {
		boot.Logger.Error("TagHandler View Execute, err =", err.Error())
		errorPage(ctx, err.Error())
	}
}

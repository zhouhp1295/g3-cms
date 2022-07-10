package page

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
)

func WriterHandler(ctx *gin.Context) {
	view, err := jetEngine().Set.GetTemplate("writer.html")
	if err != nil {
		boot.Logger.Error("WriterHandler, err = %s", err.Error())
		errorPage(ctx, err.Error())
		return
	}
	page := parseIntParam(ctx, "page", 1)
	writerId := int64(parseIntParam(ctx, "writerId", 0))
	writer := dao.ContentWriterDao.FindByPk(writerId)
	articles, pageData := dao.ContentArticleDao.FrontWriterArticles(writerId, page)

	data := globalVarMap()
	data.Set("writer", writer)
	data.Set("articles", articles)
	data.Set("pageUrl", "category")
	data.Set("pageData", pageData)
	data.Set("rightLatestHotArticles", dao.ContentArticleDao.FrontWriterRightLatestHotArticles(writerId))
	data.Set("rightRecommendArticles", dao.ContentArticleDao.FrontWriterRightRecommendArticles(writerId))

	err = view.Execute(ctx.Writer, data, nil)

	if err != nil {
		boot.Logger.Error("WriterHandler View Execute, err =", err.Error())
		errorPage(ctx, err.Error())
	}
}

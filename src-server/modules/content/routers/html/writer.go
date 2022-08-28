// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

//go:build http
// +build http

package page

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"go.uber.org/zap"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Html("/w", writerHandler)
		g3.GetGin().Html("/w.html", writerHandler)
		g3.GetGin().Html("/w/:writerId", writerHandler)
		g3.GetGin().Html("/w/:writerId/", writerHandler)
		g3.GetGin().Html("/w/:writerId/:page", writerHandler)
	})
}

func writerHandler(ctx *gin.Context) {
	view, err := jetEngine().Set.GetTemplate("writer.html")
	if err != nil {
		g3.ZL().Error("err", zap.Error(err))
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
		g3.ZL().Error("err", zap.Error(err))
		errorPage(ctx, err.Error())
	}
}

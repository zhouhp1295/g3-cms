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
		g3.GetGin().Html("/tag", tagHandler)
		g3.GetGin().Html("/tag.html", tagHandler)
		g3.GetGin().Html("/tag/:tagId", tagHandler)
		g3.GetGin().Html("/tag/:tagId/", tagHandler)
		g3.GetGin().Html("/tag/:tagId/:page", tagHandler)
	})
}

func tagHandler(ctx *gin.Context) {
	view, err := jetEngine().Set.GetTemplate("tag.html")
	if err != nil {
		g3.ZL().Error("err", zap.Error(err))
		errorPage(ctx, err.Error())
		return
	}
	page := parseIntParam(ctx, "page", 1)
	tagId := int64(parseIntParam(ctx, "tagId", 0))
	tag := dao.ContentTagDao.FindByPk(tagId)
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
		g3.ZL().Error("err", zap.Error(err))
		errorPage(ctx, err.Error())
	}
}

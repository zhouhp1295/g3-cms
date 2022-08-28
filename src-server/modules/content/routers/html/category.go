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
		g3.GetGin().Html("/category", categoryHandler)
		g3.GetGin().Html("/category.html", categoryHandler)
		g3.GetGin().Html("/category/:categoryId", categoryHandler)
		g3.GetGin().Html("/category/:categoryId/", categoryHandler)
		g3.GetGin().Html("/category/:categoryId/:page", categoryHandler)
	})
}

func categoryHandler(ctx *gin.Context) {
	view, err := jetEngine().Set.GetTemplate("category.html")
	if err != nil {
		g3.ZL().Error("err", zap.Error(err))
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
		g3.ZL().Error("err", zap.Error(err))
		errorPage(ctx, err.Error())
	}
}

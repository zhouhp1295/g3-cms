// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

//go:build http
// +build http

package page

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Html("/article", articleInfoHandler)
		g3.GetGin().Html("/article.html", articleInfoHandler)
		g3.GetGin().Html("/article/:articleId", articleInfoHandler)
	})
}

func articleInfoHandler(ctx *gin.Context) {
	pArticleId := ctx.Param("articleId")
	if len(pArticleId) == 0 {
		//todo
		return
	}
	pArticleId = strings.TrimRight(pArticleId, ".html")
	articleId, err := strconv.ParseInt(pArticleId, 10, 64)
	if err != nil {
		//todo
		return
	}
	article := dao.ContentArticleDao.FindByPk(articleId)
	if article == nil {
		//todo
		return
	}
	view, err := jetEngine().Set.GetTemplate("article.html")
	if err != nil {
		g3.ZL().Error("ArticleInfoHandler", zap.Error(err))
		errorPage(ctx, err.Error())
		return
	}
	mArticle, _ := article.(*model.ContentArticle)
	dao.ContentArticleDao.ApplyFrontData(mArticle)

	data := globalVarMap()
	data.Set("article", mArticle)
	data.Set("rightLatestHotArticles", dao.ContentArticleDao.FrontRightLatestHotArticles(mArticle.Category))
	data.Set("rightRecommendArticles", dao.ContentArticleDao.FrontRightRecommendArticles(mArticle.Category))

	err = view.Execute(ctx.Writer, data, nil)
	if err != nil {
		g3.ZL().Error("IndexHandler View Execute", zap.Error(err))
		errorPage(ctx, err.Error())
	}
}

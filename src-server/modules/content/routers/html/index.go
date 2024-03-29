// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

//go:build http
// +build http

package page

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3-cms/modules/render"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var _jetEngine *render.JetEngine

var jetEngineOnce sync.Once

func init() {
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Html("/robots.txt", robotsHandler)
	})
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Html("/", indexHandler)
		g3.GetGin().Html("/index", indexHandler)
		g3.GetGin().Html("/index.html", indexHandler)

		g3.GetGin().Html("/latest", latestHandler)
		g3.GetGin().Html("/latest.html", latestHandler)
		g3.GetGin().Html("/latest/:page", latestHandler)
	})
}

func robotsHandler(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(dao.ContentConfigDao.WebConfig().Robots))
}

func jetEngine() *render.JetEngine {
	jetEngineOnce.Do(func() {
		cfg := dao.ContentConfigDao.WebConfig()
		host := ""
		if len(cfg.Host) > 0 {
			host = cfg.Host
		}
		_jetEngine = render.New(host, "default")

		_jetEngine.Set.SetDevelopmentMode(true)

		addJetFunc(_jetEngine.Host, _jetEngine.Set)
	})
	return _jetEngine
}

func addJetFunc(host string, s *jet.Set) {

	s.AddGlobalFunc("writerUrl", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("writerUrl", 1, 1)
		v := a.Get(0)
		if v.Kind() != reflect.Int64 {
			return v
		}
		return reflect.ValueOf(fmt.Sprintf("%s/w/%d.html", host, v.Int()))
	})

	s.AddGlobalFunc("articleUrl", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("articleUrl", 1, 1)
		return reflect.ValueOf(fmt.Sprintf("%s/article/%s.html", host, render.ArgToString(a.Get(0))))
	})

	s.AddGlobalFunc("tagUrl", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("tagUrl", 1, 1)
		return reflect.ValueOf(fmt.Sprintf("%s/tag/%s.html", host, render.ArgToString(a.Get(0))))
	})

	s.AddGlobalFunc("categoryUrl", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("cmsCategoryUrl", 1, 1)
		return reflect.ValueOf(fmt.Sprintf("%s/category/%s.html", host, render.ArgToString(a.Get(0))))
	})
}

func globalVarMap() jet.VarMap {
	varMap := make(jet.VarMap)
	varMap.Set("title", "")
	varMap.Set("keywords", "")
	varMap.Set("description", "")
	varMap.Set("cfg", dao.ContentConfigDao.WebConfig())
	varMap.Set("menus", dao.ContentMenuDao.FrontMenus())
	varMap.Set("rightTags", dao.ContentTagDao.FrontTags())
	return varMap
}

func errorPage(ctx *gin.Context, msg ...string) {

}

func parseParam(ctx *gin.Context, field string) string {
	param := ctx.Param(field)
	if len(param) == 0 {
		return param
	}
	param = strings.TrimLeft(param, "/")
	param = strings.Replace(param, ".html", "", 1)
	return param
}

func parseIntParam(ctx *gin.Context, field string, min int) int {
	if min < 0 {
		min = 0
	}
	param := ctx.Param(field)
	if len(param) == 0 {
		return min
	}
	param = strings.TrimLeft(param, "/")
	param = strings.Replace(param, ".html", "", 1)
	i, err := strconv.Atoi(param)
	if err != nil {
		g3.ZL().Error("err", zap.String("param", param), zap.Error(err))
		return min
	}
	return i
}

func indexHandler(ctx *gin.Context) {
	view, err := jetEngine().Set.GetTemplate("index.html")
	if err != nil {
		g3.ZL().Error("err", zap.Error(err))
		errorPage(ctx, err.Error())
		return
	}

	data := globalVarMap()
	data.Set("banners", dao.ContentBannerDao.FrontBanners())
	data.Set("topArticles", dao.ContentArticleDao.FrontTopArticles())
	latestArticles, _ := dao.ContentArticleDao.FrontLatestArticles(1)
	data.Set("latestArticles", latestArticles)
	data.Set("rightLatestHotArticles", dao.ContentArticleDao.FrontRightLatestHotArticles(0))
	data.Set("rightRecommendArticles", dao.ContentArticleDao.FrontRightRecommendArticles(0))

	err = view.Execute(ctx.Writer, data, nil)
	if err != nil {
		g3.ZL().Error("err", zap.Error(err))
		errorPage(ctx, err.Error())
	}
}

func latestHandler(ctx *gin.Context) {
	view, err := jetEngine().Set.GetTemplate("latest.html")
	if err != nil {
		g3.ZL().Error("err", zap.Error(err))
		errorPage(ctx, err.Error())
		return
	}
	page := parseIntParam(ctx, "page", 1)

	articles, pageData := dao.ContentArticleDao.FrontLatestArticles(page)

	data := globalVarMap()
	data.Set("articles", articles)
	data.Set("pageUrl", "latest")
	data.Set("pageData", pageData)
	data.Set("rightLatestHotArticles", dao.ContentArticleDao.FrontRightLatestHotArticles(0))
	data.Set("rightRecommendArticles", dao.ContentArticleDao.FrontRightRecommendArticles(0))

	err = view.Execute(ctx.Writer, data, nil)

	if err != nil {
		g3.ZL().Error("err", zap.Error(err))
		errorPage(ctx, err.Error())
	}
}

package page

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/dao"
	"github.com/zhouhp1295/g3-cms/render"
	"reflect"
	"sync"
)

var _jetEngine *render.JetEngine

var jetEngineOnce sync.Once

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
		v := a.Get(0)
		if v.Kind() != reflect.Int64 {
			return v
		}
		return reflect.ValueOf(fmt.Sprintf("%s/article/%d.html", host, v.Int()))
	})

	s.AddGlobalFunc("tagUrl", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("tagUrl", 1, 1)
		v := a.Get(0)
		if v.Kind() != reflect.Int64 {
			return v
		}
		return reflect.ValueOf(fmt.Sprintf("%s/tag/%d.html", host, v.Int()))
	})

	s.AddGlobalFunc("categoryUrl", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("cmsCategoryUrl", 1, 1)
		v := a.Get(0)
		if v.Kind() != reflect.Int64 {
			return v
		}
		return reflect.ValueOf(fmt.Sprintf("%s/category/%d.html", host, v.Int()))
	})
}

func globalVarMap() jet.VarMap {
	varMap := make(jet.VarMap)
	varMap.Set("title", "")
	varMap.Set("keywords", "")
	varMap.Set("description", "")
	varMap.Set("cfg", dao.ContentConfigDao.WebConfig())
	return varMap
}

func IndexHandler(ctx *gin.Context) {
	view, err := jetEngine().Set.GetTemplate("index.html")
	if err != nil {
		boot.Logger.Error("IndexHandler, err = %s", err.Error())
		errorPage(ctx, err.Error())
		return
	}

	data := globalVarMap()
	data.Set("banners", dao.ContentBannerDao.FrontBanners())
	data.Set("topArticles", dao.ContentArticleDao.FrontTopArticles())
	data.Set("menus", dao.ContentMenuDao.FrontMenus())
	latestArticles, _ := dao.ContentArticleDao.FrontLatestArticles(1)
	data.Set("latestArticles", latestArticles)

	err = view.Execute(ctx.Writer, data, nil)
	if err != nil {
		boot.Logger.Error("IndexHandler View Execute, err =", err.Error())
		errorPage(ctx, err.Error())
	}
}

func LatestHandler(ctx *gin.Context) {

}

func errorPage(ctx *gin.Context, msg ...string) {

}

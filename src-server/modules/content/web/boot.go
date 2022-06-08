package web

import (
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/web/page"
)

func initWebPageRouters() {
	boot.WebPage("/", page.IndexHandler)
	boot.WebPage("/index", page.IndexHandler)
	boot.WebPage("/index.html", page.IndexHandler)
	boot.WebPage("/article", page.ArticleInfoHandler)
	boot.WebPage("/article.html", page.ArticleInfoHandler)
	boot.WebPage("/article/:articleId", page.ArticleInfoHandler)
	boot.WebPage("/latest", page.LatestHandler)
	boot.WebPage("/latest.html", page.LatestHandler)
	boot.WebPage("/latest/:page", page.LatestHandler)
	boot.WebPage("/category", page.CategoryHandler)
	boot.WebPage("/category.html", page.CategoryHandler)
	boot.WebPage("/category/:categoryId", page.CategoryHandler)
	boot.WebPage("/category/:categoryId/", page.CategoryHandler)
	boot.WebPage("/category/:categoryId/:page", page.CategoryHandler)
	boot.WebPage("/tag", page.TagHandler)
	boot.WebPage("/tag.html", page.TagHandler)
	boot.WebPage("/tag/:tagId", page.TagHandler)
	boot.WebPage("/tag/:tagId/", page.TagHandler)
	boot.WebPage("/tag/:tagId/:page", page.TagHandler)
	boot.WebPage("/w", page.WriterHandler)
	boot.WebPage("/w.html", page.WriterHandler)
	boot.WebPage("/w/:writerId", page.WriterHandler)
	boot.WebPage("/w/:writerId/", page.WriterHandler)
	boot.WebPage("/w/:writerId/:page", page.WriterHandler)
}

func initWebApiRouters() {

}

func Init() {
	initWebPageRouters()
}

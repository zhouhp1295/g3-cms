package dao

import (
	"fmt"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
	"sort"
	"time"
)

func clearFrontAllBannerCache() {
	boot.Lache.Delete("K-Content-Dao-Banner-FrontBanners")
}

func getFrontAllBanners() (banners []FrontBannerData, ok bool) {
	query := crud.DbSess()
	banners = make([]FrontBannerData, 0)
	webConfig := getWebConfigFromCache()
	//Step 1 获取所有的Banner
	bannerRows := make([]model.ContentBanner, 0)
	query.Where("status = ? and deleted = ?", crud.FlagYes, crud.FlagNo).Find(&bannerRows)
	for _, bannerRow := range bannerRows {
		banners = append(banners, FrontBannerData{
			Title: bannerRow.Title,
			Url:   bannerRow.Url,
			Pic:   bannerRow.Cover,
			Sort:  bannerRow.Sort,
		})
	}
	//Step 2 获取所有显示到Banner的分类
	categoryRows := make([]model.ContentCategory, 0)
	query.Where("status = ? and in_banner = ? and deleted = ? ", crud.FlagYes, crud.FlagYes, crud.FlagNo).Find(&categoryRows)
	for _, categoryRow := range categoryRows {
		pic := categoryRow.Banner
		if len(pic) == 0 {
			pic = categoryRow.Cover
		}
		banners = append(banners, FrontBannerData{
			Title: categoryRow.Title,
			Url:   fmt.Sprintf("%s/category/%d.html", webConfig.Host, categoryRow.Id),
			Pic:   pic,
			Sort:  categoryRow.InBannerSort,
		})
	}
	//Step 3 获取所有显示到Banner的文章
	articleRows := make([]model.ContentArticle, 0)
	query.Where("status = ? and in_banner = ? and deleted = ? ", crud.FlagYes, crud.FlagYes, crud.FlagNo).Find(&categoryRows)
	for _, articleRow := range articleRows {
		pic := articleRow.Banner
		if len(pic) == 0 {
			pic = articleRow.Cover
		}
		banners = append(banners, FrontBannerData{
			Title: articleRow.Title,
			Url:   fmt.Sprintf("%s/article/%d.html", webConfig.Host, articleRow.Id),
			Pic:   pic,
			Sort:  articleRow.InBannerSort,
		})
	}
	//Step 4 排序
	sort.Slice(banners, func(i, j int) bool {
		return banners[i].Sort < banners[j].Sort
	})
	return
}

func getFrontAllBannersFromCache() []FrontBannerData {
	key := "K-Content-Dao-Banner-FrontBanners"
	result := make([]FrontBannerData, 0)
	ok := boot.Lache.GetT(key, &result)
	if !ok {
		_data, _ok := getFrontAllBanners()
		if _ok {
			boot.Lache.Set(key, _data, 30*time.Minute)
		} else {
			boot.Lache.Set(key, _data, 10*time.Second)
		}
		return _data
	}
	return result
}

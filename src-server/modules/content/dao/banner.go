package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
)

type contentBannerDAO struct {
	crud.BaseDao
}

var ContentBannerDao = &contentBannerDAO{
	crud.BaseDao{Model: new(model.ContentBanner)},
}

type FrontBannerData struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Pic   string `json:"pic"`
	Sort  int    `json:"sort"`
}

func (dao *contentBannerDAO) FrontBanners() []FrontBannerData {
	return getFrontAllBannersFromCache()
}

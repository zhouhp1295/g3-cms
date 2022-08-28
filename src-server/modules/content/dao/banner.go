// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

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

func (dao *contentBannerDAO) ClearCache() {
	clearFrontAllBannerCache()
}
func (dao *contentBannerDAO) AfterInsert(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	dao.ClearCache()
	return
}

func (dao *contentBannerDAO) AfterUpdate(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	dao.ClearCache()
	return
}

func (dao *contentBannerDAO) AfterDelete(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	dao.ClearCache()
	return
}

func (dao *contentBannerDAO) AfterRemove(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	dao.ClearCache()
	return
}

func (dao *contentBannerDAO) UpdateStatus(pk int64, status interface{}, operator int64) bool {
	m := dao.FindByPk(pk)
	_m, _ := m.(*model.ContentBanner)
	_status, _ := status.(string)
	_m.Status = _status
	dao.Update(_m, operator)
	dao.ClearCache()
	return true
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

package dao

import (
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
	"github.com/zhouhp1295/lache/driver"
	"time"
)

func clearFrontAllCategoryCache() {
	boot.Lache.Delete("K-Content-Dao-Category-FrontTreeOptions")
}

func listFrontCategoryTreeOptions() []helpers.TreeOption {
	search := &model.ContentCategory{
		TailColumns: crud.TailColumns{
			Status:  crud.FlagYes,
			Deleted: crud.FlagNo,
		},
	}
	allRows := ContentCategoryDao.FindAll(search, nil)
	rows, ok := allRows.([]model.ContentCategory)
	if !ok {
		return []helpers.TreeOption{}
	}
	result := make([]helpers.TreeOption, 0)
	for _, row := range rows {
		result = append(result, helpers.TreeOption{
			Id:    row.Id,
			Pid:   row.Pid,
			Label: row.Title,
		})
	}
	return result
}

func listFrontCategoryTreeOptionsFromCache() []helpers.TreeOption {
	key := "K-Content-Dao-Category-FrontTreeOptions"
	var result []helpers.TreeOption
	ok := boot.Lache.GetT(key, &result)
	if !ok {
		options := listFrontCategoryTreeOptions()
		if len(options) == 0 {
			boot.Lache.Set(key, options, 10*time.Second)
		} else {
			boot.Lache.Set(key, options, driver.NotExpired)
		}
		return options
	}
	return result
}

func getCategoryName(categoryId int64) string {
	if categoryId <= 0 {
		return "未分类"
	}
	options := listFrontCategoryTreeOptionsFromCache()
	for _, option := range options {
		if option.Id == categoryId {
			return option.Label
		}
	}
	return "未分类"
}

func getCategoryLevels(categoryId int64) []helpers.TreeOption {
	return helpers.ToParentLevels(listFrontCategoryTreeOptionsFromCache(), categoryId)
}

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
	boot.Lache.Delete("K-Content-Dao-Category-FrontOptions")
}

func listFrontCategorySelectOptions() []helpers.SelectOption {
	search := &model.ContentCategory{
		TailColumns: crud.TailColumns{
			Status:  crud.FlagYes,
			Deleted: crud.FlagNo,
		},
	}
	allRows := ContentWriterDao.FindAll(search, nil)
	rows, ok := allRows.([]model.ContentWriter)
	if !ok {
		return []helpers.SelectOption{}
	}
	result := make([]helpers.SelectOption, 0)
	for _, row := range rows {
		result = append(result, helpers.SelectOption{
			Id:    row.Id,
			Label: row.Name,
		})
	}
	return result
}

func listCategorySelectOptionsFromCache() []helpers.SelectOption {
	key := "K-Content-Dao-Category-FrontOptions"
	var result []helpers.SelectOption
	ok := boot.Lache.GetT(key, &result)
	if !ok {
		options := listFrontCategorySelectOptions()
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
	options := listCategorySelectOptionsFromCache()
	for _, option := range options {
		if option.Id == categoryId {
			return option.Label
		}
	}
	return "未分类"
}

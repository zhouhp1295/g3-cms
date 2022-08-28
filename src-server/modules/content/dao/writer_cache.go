// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package dao

import (
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
	"github.com/zhouhp1295/lache/driver"
	"time"
)

func clearAllWriterCache() {
	boot.Lache.Delete("K-Content-Dao-Writer-SelectOptions")
}

func listWriterSelectOptions() []helpers.SelectOption {
	searchRole := &model.ContentWriter{
		TailColumns: crud.TailColumns{
			Status:  crud.FlagYes,
			Deleted: crud.FlagNo,
		},
	}
	allRoleData := ContentWriterDao.FindAll(searchRole, nil)
	roles, ok := allRoleData.([]model.ContentWriter)
	if !ok {
		return []helpers.SelectOption{}
	}
	result := make([]helpers.SelectOption, 0)
	for _, row := range roles {
		result = append(result, helpers.SelectOption{
			Id:    row.Id,
			Label: row.Name,
		})
	}
	return result
}

func listWriterSelectOptionsFromCache() []helpers.SelectOption {
	key := "K-Content-Dao-Writer-SelectOptions"
	var result []helpers.SelectOption
	ok := boot.Lache.GetT(key, &result)
	if !ok {
		options := listWriterSelectOptions()
		if len(options) == 0 {
			boot.Lache.Set(key, options, 10*time.Second)
		} else {
			boot.Lache.Set(key, options, driver.NotExpired)
		}
		return options
	}
	return result
}

func getWriterName(writerId int64) string {
	if writerId <= 0 {
		return "佚名"
	}
	options := listWriterSelectOptionsFromCache()
	for _, option := range options {
		if option.Id == writerId {
			return option.Label
		}
	}
	return ""
}

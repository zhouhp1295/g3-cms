// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package dao

import (
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
	"github.com/zhouhp1295/lache/driver"
	"time"
)

func clearAllMenuCache() {
	boot.Lache.Delete("K-System-Dao-Menu-AllVisible")
	boot.Lache.Delete("K-System-Dao-Menu-TreeOptions")
}

func listAllVisibleMenus() []model.SysMenu {
	key := "K-System-Dao-Menu-AllVisible"
	var result []model.SysMenu
	ok := boot.Lache.GetT(key, &result)
	if !ok {
		searchMenu := &model.SysMenu{
			Visible: crud.FlagYes,
			TailColumns: crud.TailColumns{
				Status:  crud.FlagYes,
				Deleted: crud.FlagNo,
			},
		}
		allMenus := SysMenuDao.FindAll(searchMenu, &crud.BaseQueryParams{
			OrderBy: "sort ASC",
		})
		rows, ok := allMenus.([]model.SysMenu)
		if !ok {
			return result
		}
		if len(rows) == 0 {
			boot.Lache.Set(key, rows, 10*time.Second)
		} else {
			boot.Lache.Set(key, rows, driver.NotExpired)
		}
		return rows
	}
	return result
}

func listMenuTree() []helpers.TreeNode {
	allMenus := listAllVisibleMenus()
	if len(allMenus) == 0 {
		return []helpers.TreeNode{}
	}
	items := make([]helpers.TreeOption, len(allMenus))
	for i := 0; i < len(items); i++ {
		items[i] = helpers.TreeOption{
			Id:    allMenus[i].Id,
			Pid:   allMenus[i].Pid,
			Label: allMenus[i].Title,
		}
	}
	return helpers.ToTree(items)
}

func listMenuTreeFromCache() []helpers.TreeNode {
	key := "K-System-Dao-Menu-TreeOptions"
	var result []helpers.TreeNode
	ok := boot.Lache.GetT(key, &result)
	if !ok {
		options := listMenuTree()
		if len(options) == 0 {
			boot.Lache.Set(key, options, 10*time.Second)
		} else {
			boot.Lache.Set(key, options, driver.NotExpired)
		}
		return options
	}
	return result
}

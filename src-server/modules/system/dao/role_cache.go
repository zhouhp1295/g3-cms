// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package dao

import (
	"fmt"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
	"github.com/zhouhp1295/lache/driver"
	"strings"
	"time"
)

type CachedRoleData struct {
	Identifier string
	Perms      []string
}

func clearAllRoleCache() {
	boot.Lache.Delete("K-System-Dao-Role-Perms")
	boot.Lache.Delete("K-System-Dao-Role-SelectOptions")
}

func listRolePermsFromCache() map[string]CachedRoleData {
	key := "K-System-Dao-Role-Perms"
	var result map[string]CachedRoleData
	ok := boot.Lache.GetT(key, &result)
	if !ok {
		permsData := listRolePerms()
		if len(permsData) == 0 {
			boot.Lache.Set(key, permsData, 10*time.Second)
		} else {
			boot.Lache.Set(key, permsData, driver.NotExpired)
		}
		return permsData
	}
	return result
}

func setPerms(roleData map[string]CachedRoleData) {
	rg := g3.GetGin().Group("/api")
	rg.ClearAllRolesPerm()
	for _, data := range roleData {
		rg.AddRolePerm(data.Identifier, data.Perms...)
	}
}

func listRolePerms() map[string]CachedRoleData {
	result := make(map[string]CachedRoleData)
	searchRole := &model.SysRole{
		TailColumns: crud.TailColumns{
			Status:  crud.FlagYes,
			Deleted: crud.FlagNo,
		},
	}
	//取所有的角色
	allRoleData := SysRoleDao.FindAll(searchRole, nil)
	roles, ok := allRoleData.([]model.SysRole)
	if !ok {
		return result
	}
	//取所有的菜单
	menus := SysMenuDao.FindAllVisibleMenus()
	//list=>map
	menuMap := make(map[string]model.SysMenu)
	for _, menu := range menus {
		menuMap[fmt.Sprintf("%d", menu.Id)] = menu
	}
	//组织每个角色对应的权限
	for _, role := range roles {
		roleId := fmt.Sprintf("%d", role.Id)
		data := CachedRoleData{
			Identifier: role.Identifier,
			Perms:      make([]string, 0),
		}
		roleMenus := strings.Split(role.Menus, ",")
		for _, roleMenu := range roleMenus {
			if menu, exist := menuMap[roleMenu]; exist {
				data.Perms = append(data.Perms, menu.Perms)
			}
		}
		result[roleId] = data
	}
	return result
}

func listRoleOptions() []helpers.SelectOption {
	searchRole := &model.SysRole{
		TailColumns: crud.TailColumns{
			Status:  crud.FlagYes,
			Deleted: crud.FlagNo,
		},
	}
	allRoleData := SysRoleDao.FindAll(searchRole, nil)
	roles, ok := allRoleData.([]model.SysRole)
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

func listRoleOptionsFromCache() []helpers.SelectOption {
	key := "K-System-Dao-Role-SelectOptions"
	var result []helpers.SelectOption
	ok := boot.Lache.GetT(key, &result)
	if !ok {
		options := listRoleOptions()
		if len(options) == 0 {
			boot.Lache.Set(key, options, 10*time.Second)
		} else {
			boot.Lache.Set(key, options, driver.NotExpired)
		}
		return options
	}
	return result
}

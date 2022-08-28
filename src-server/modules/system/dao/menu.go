// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3-cms/utils"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
)

type sysMenuDAO struct {
	crud.BaseDao
}

var SysMenuDao = &sysMenuDAO{
	crud.BaseDao{Model: new(model.SysMenu)},
}

type MenuRouterMeta struct {
	Icon    string `json:"icon"`
	Title   string `json:"title"`
	Link    string `json:"link"`
	NoCache bool   `json:"noCache"`
}

type MenuRouterData struct {
	Name       string           `json:"name"`
	Path       string           `json:"path"`
	Component  string           `json:"component"`
	Hidden     bool             `json:"hidden"`
	AlwaysShow bool             `json:"alwaysShow"`
	Redirect   string           `json:"redirect"`
	Children   []MenuRouterData `json:"children"`
	Query      string           `json:"query"`
	Meta       MenuRouterMeta   `json:"meta"`
}

func (dao *sysMenuDAO) MenuTree() []helpers.TreeNode {
	return listMenuTreeFromCache()
}

func (dao *sysMenuDAO) FindAllVisibleMenus() []model.SysMenu {
	return listAllVisibleMenus()
}

func (dao *sysMenuDAO) getMenus(menus []model.SysMenu, pid int64) (data []MenuRouterData) {
	data = make([]MenuRouterData, 0)
	for _, menu := range menus {
		if len(menu.Path) == 0 || len(menu.Name) == 0 || len(menu.Component) == 0 {
			continue
		}
		if menu.Pid == pid {
			routerData := MenuRouterData{
				Name:      menu.Name,
				Path:      menu.Path,
				Component: menu.Component,
				Hidden:    !utils.YesOrNo(menu.Visible),
				Meta: MenuRouterMeta{
					Icon:    menu.Icon,
					Title:   menu.Title,
					NoCache: utils.YesOrNo(menu.IsCache),
				},
				Children: dao.getMenus(menus, menu.Id),
			}
			data = append(data, routerData)
		}
	}
	return
}

// GetMenuRouters 获取所有的菜单
func (dao *sysMenuDAO) GetMenuRouters() (data []MenuRouterData) {
	data = make([]MenuRouterData, 0)
	menus := SysMenuDao.FindAllVisibleMenus()
	if len(menus) == 0 {
		return
	}
	// 递归
	for _, menu := range menus {
		if menu.Pid == 0 {
			routerData := MenuRouterData{
				Name:       menu.Name,
				Path:       menu.Path,
				AlwaysShow: true,
				Component:  menu.Component,
				Redirect:   "noRedirect",
				Hidden:     !utils.YesOrNo(menu.Visible),
				Meta: MenuRouterMeta{
					Icon:    menu.Icon,
					Title:   menu.Title,
					NoCache: utils.YesOrNo(menu.IsCache),
				},
				Children: dao.getMenus(menus, menu.Id),
			}
			data = append(data, routerData)
		}
	}
	return
}

func (dao *sysMenuDAO) getMenusByPerms(permissions []string, menus []model.SysMenu, pid int64) (data []MenuRouterData) {
	data = make([]MenuRouterData, 0)
	for _, menu := range menus {
		if len(menu.Perms) > 0 && helpers.IndexOf[string](permissions, menu.Perms) < 0 {
			continue
		}
		if len(menu.Path) == 0 || len(menu.Name) == 0 || len(menu.Component) == 0 {
			continue
		}
		if menu.Pid == pid {
			routerData := MenuRouterData{
				Name:      menu.Name,
				Path:      menu.Path,
				Component: menu.Component,
				Hidden:    !utils.YesOrNo(menu.Visible),
				Meta: MenuRouterMeta{
					Icon:    menu.Icon,
					Title:   menu.Title,
					NoCache: utils.YesOrNo(menu.IsCache),
				},
				Children: dao.getMenusByPerms(permissions, menus, menu.Id),
			}
			data = append(data, routerData)
		}
	}
	return
}

// GetMenuRoutersByPerms 根据权限获取可访问的菜单
func (dao *sysMenuDAO) GetMenuRoutersByPerms(permissions []string) (data []MenuRouterData) {
	data = make([]MenuRouterData, 0)
	menus := SysMenuDao.FindAllVisibleMenus()
	if len(menus) == 0 {
		return
	}
	//嵌套处理
	for _, menu := range menus {
		if len(menu.Perms) > 0 && helpers.IndexOf[string](permissions, menu.Perms) < 0 {
			continue
		}
		if menu.Pid == 0 {
			routerData := MenuRouterData{
				Name:       menu.Name,
				Path:       menu.Path,
				AlwaysShow: true,
				Component:  menu.Component,
				Redirect:   "noRedirect",
				Hidden:     !utils.YesOrNo(menu.Visible),
				Meta: MenuRouterMeta{
					Icon:    menu.Icon,
					Title:   menu.Title,
					NoCache: utils.YesOrNo(menu.IsCache),
				},
				Children: dao.getMenusByPerms(permissions, menus, menu.Id),
			}
			if len(routerData.Children) > 0 {
				data = append(data, routerData)
			}
		}
	}
	return
}

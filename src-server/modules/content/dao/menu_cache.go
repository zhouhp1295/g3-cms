package dao

import (
	"fmt"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
	"sort"
	"time"
)

func clearFrontAllMenuCache() {
	boot.Lache.Delete("K-Content-Dao-Menu-FrontMenus")
}

func getFrontAllMenus() (menus []FrontMenuData, ok bool) {
	query := crud.DbSess()
	menus = make([]FrontMenuData, 0)
	webConfig := getWebConfigFromCache()
	//Step 1 获取所有的Menu
	menuRows := make([]model.ContentMenu, 0)
	query.Where("status = ? and delete = ?", crud.FlagYes, crud.FlagNo).Find(&menuRows)
	for _, menuRow := range menuRows {
		menus = append(menus, FrontMenuData{
			Title: menuRow.Title,
			Url:   menuRow.Url,
			Icon:  menuRow.Icon,
			Sort:  menuRow.Sort,
		})
	}
	//Step 2 获取所有显示到Menu的分类
	categoryRows := make([]model.ContentCategory, 0)
	query.Where("status = ? and in_menu = ? and deleted = ? ", crud.FlagYes, crud.FlagYes, crud.FlagNo).Find(&categoryRows)
	for _, categoryRow := range categoryRows {
		icon := categoryRow.Icon
		if len(icon) == 0 {
			icon = categoryRow.Cover
		}
		menus = append(menus, FrontMenuData{
			Title: categoryRow.Title,
			Url:   fmt.Sprintf("%s/category/%d.html", webConfig.Host, categoryRow.Id),
			Icon:  icon,
			Sort:  categoryRow.InMenuSort,
		})
	}
	//Step 3 排序
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Sort < menus[j].Sort
	})
	return
}

func getFrontAllMenusFromCache() []FrontMenuData {
	key := "K-Content-Dao-Menu-FrontMenus"
	result := make([]FrontMenuData, 0)
	ok := boot.Lache.GetT(key, &result)
	if !ok {
		_data, _ok := getFrontAllMenus()
		if _ok {
			boot.Lache.Set(key, _data, 30*time.Minute)
		} else {
			boot.Lache.Set(key, _data, 10*time.Second)
		}
		return _data
	}
	return result
}

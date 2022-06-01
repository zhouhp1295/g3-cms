package dao

import (
	"fmt"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
	"strings"
)

type contentCategoryDAO struct {
	crud.BaseDao
}

var ContentCategoryDao = &contentCategoryDAO{
	crud.BaseDao{Model: new(model.ContentCategory)},
}

func (dao *contentCategoryDAO) BeforeInsert(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.ContentCategory); _ok {
		//祖级列表
		if _m.Pid > 0 {
			cnt := dao.CountByPk(_m.Pid)
			if cnt == 0 {
				msg = "父级分类不存在"
				return
			}
			existM3 := dao.FindByPk(_m.Pid)
			_m3, _ := existM3.(*model.ContentCategory)
			_m.Ancestors = fmt.Sprintf("%s,%d,", _m3.Ancestors, _m.Pid)
		} else {
			_m.Ancestors = "0,"
		}
		ok = true
	}
	return
}

func (dao *contentCategoryDAO) BeforeUpdate(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.ContentCategory); _ok {
		if _m.Id == _m.Pid {
			msg = "无法设置为自己的子级分类"
			return
		}
		exist := dao.FindByPk(_m.Id)
		_m.SetLastModel(exist)
		_m2, _ := exist.(*model.ContentCategory)
		if _m.Pid != _m2.Pid {
			//父级变化后，需要调整祖级列表
			if _m.Pid > 0 {
				cnt := dao.CountByPk(_m.Pid)
				if cnt == 0 {
					msg = "父级分类不存在"
					return
				}
				existM3 := dao.FindByPk(_m.Pid)
				_m3, _ := existM3.(*model.ContentCategory)
				if strings.Contains(_m3.Ancestors, fmt.Sprintf(",%d,", _m.Id)) {
					msg = "无法设置子级分类为自己的父级分类"
					return
				}
				_m.Ancestors = fmt.Sprintf("%s%d,", _m3.Ancestors, _m.Pid)
			} else {
				_m.Ancestors = "0,"
			}
		}
		ok = true
	}
	return
}

func (dao *contentCategoryDAO) updateChildrenAncestors(m *model.ContentCategory) {
	children := make([]model.ContentCategory, 0)
	if crud.DbSess().Where("pid = ? and deleted = ?", m.Id, crud.FlagNo).Find(&children).Error == nil {
		for _, m2 := range children {
			m2.Ancestors = fmt.Sprintf("%s%d,", m.Ancestors, m2.Pid)
			crud.DbSess().Select([]string{"ancestors"}).Updates(m2)
			dao.updateChildrenAncestors(&m2)
		}
	}
}

func (dao *contentCategoryDAO) AfterUpdate(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.ContentCategory); _ok {
		if _m.Last != nil {
			_m2, _ := _m.Last.(*model.ContentCategory)
			if _m.Pid != _m2.Pid {
				//更新子级分类的祖级列表
				dao.updateChildrenAncestors(_m)
			}
		}
		ok = true
	}
	return
}

func (dao *contentCategoryDAO) BeforeDelete(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.ContentCategory); _ok {
		if boot.UseMysql {
			if dao.Count(" find_in_set (?, ancestors)", _m.Id) > 0 {
				msg = "存在子分类"
				return
			}
		} else {
			if dao.Count(" ancestors like ?", fmt.Sprintf("%%,%d,%%", _m.Id)) > 0 {
				msg = "存在子分类"
				return
			}
		}
		ok = true
	}
	return
}

func (dao *contentCategoryDAO) BeforeRemove(m crud.ModelInterface) (ok bool, msg string) {
	ok, msg = dao.BeforeDelete(m)
	return
}

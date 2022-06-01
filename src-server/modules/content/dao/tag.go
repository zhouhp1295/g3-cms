package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
)

type contentTagDAO struct {
	crud.BaseDao
}

var ContentTagDao = &contentTagDAO{
	crud.BaseDao{Model: new(model.ContentTag)},
}

func (dao *contentTagDAO) InsertOrGetByTitle(title string) *model.ContentTag {
	cnt := dao.CountByColumn("title", title)
	m := new(model.ContentTag)
	if cnt > 0 {
		//已存在，取并返回
		if crud.DbSess().Where("title = ? and deleted = ?", title, crud.FlagNo).First(m).Error == nil {
			return m
		}
	} else {
		//新的标签
		m.Title = title
		if crud.DbSess().Create(m).Error == nil {
			return m
		}
	}
	return nil
}

func (dao *contentTagDAO) BeforeInsert(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.ContentTag); _ok {
		cnt := dao.CountByColumn("title", _m.Title)
		if cnt > 0 {
			msg = "已存在标签名"
			return
		}
		ok = true
	}
	return
}

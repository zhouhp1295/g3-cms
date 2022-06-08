package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/content/model"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
)

type contentWriterDAO struct {
	crud.BaseDao
}

var ContentWriterDao = &contentWriterDAO{
	crud.BaseDao{Model: new(model.ContentWriter)},
}

func (dao *contentWriterDAO) GetWriterName(pk int64) string {
	return getWriterName(pk)
}

func (dao *contentWriterDAO) SelectOptions() []helpers.SelectOption {
	return listWriterSelectOptionsFromCache()
}

func (dao *contentWriterDAO) AfterInsert(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	clearAllWriterCache()
	return
}
func (dao *contentWriterDAO) AfterUpdate(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	clearAllWriterCache()
	return
}
func (dao *contentWriterDAO) AfterDelete(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	clearAllWriterCache()
	return
}

func (dao *contentWriterDAO) AfterRemove(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	clearAllWriterCache()
	return
}

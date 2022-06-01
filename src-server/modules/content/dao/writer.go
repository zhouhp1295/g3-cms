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

func (dao *contentWriterDAO) SelectOptions() []helpers.SelectOption {
	return listWriterSelectOptionsFromCache()
}

package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/crud"
)

type sysDictTypeDAO struct {
	crud.BaseDao
}

var SysDictTypeDao = &sysDictTypeDAO{
	crud.BaseDao{Model: new(model.SysDictType)},
}

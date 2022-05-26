package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/crud"
)

type sysDictDataDAO struct {
	crud.BaseDao
}

var SysDictDataDao = &sysDictDataDAO{
	crud.BaseDao{Model: new(model.SysDictData)},
}

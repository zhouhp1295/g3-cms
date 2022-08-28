// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

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

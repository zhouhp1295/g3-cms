// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/crud"
)

type sysConfigDAO struct {
	crud.BaseDao
}

var SysConfigDao = &sysConfigDAO{
	crud.BaseDao{Model: new(model.SysConfig)},
}

func (dao *sysConfigDAO) BeforeInsert(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.SysConfig); _ok {
		if dao.CountByColumn("code", _m.Code) > 0 {
			msg = "Code已存在"
			return
		}
		ok = true
	}
	return
}

func (dao *sysConfigDAO) BeforeUpdate(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.SysConfig); _ok {
		if dao.CountByColumn("code", _m.Code) > 1 {
			msg = "Code已存在"
			return
		}
		ok = true
	}
	return
}

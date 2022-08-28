// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package dao

import (
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
	"go.uber.org/zap"
	"gorm.io/gorm/utils"
)

const SuperUserID = 999999

type sysUserDAO struct {
	crud.BaseDao
}

var SysUserDao = &sysUserDAO{
	crud.BaseDao{Model: new(model.SysUser)},
}

func (dao *sysUserDAO) CreateSuperUser(username, password string) bool {
	existNum := dao.CountByPk(SuperUserID)
	if existNum == 0 {
		hashedPwd, err := helpers.PasswordHash(password)
		if err != nil {
			g3.ZL().Fatal("createSuperUser Failed", zap.Error(err))
			return false
		}
		superUser := new(model.SysUser)
		superUser.Id = SuperUserID
		superUser.Username = username
		superUser.Password = hashedPwd
		superUser.Nickname = "超级管理员"
		superUser.Roles = utils.ToString(SuperRoleID)
		err = crud.DbSess().Create(superUser).Error
		if err != nil {
			g3.ZL().Fatal("createSuperUser Failed", zap.Error(err))
			return false
		}
		return true
	}
	return false
}

func (dao *sysUserDAO) BeforeInsert(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.SysUser); _ok {
		if dao.CountByColumn("username", _m.Username) > 0 {
			msg = "用户名已存在"
			return
		}
		_m.Password, _ = helpers.PasswordHash(_m.Password)
		ok = true
	}
	return
}

func (dao *sysUserDAO) BeforeUpdate(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.SysUser); _ok {
		if dao.CountByColumn("username", _m.Username) > 1 {
			msg = "用户名已存在"
			return
		}
		ok = true
	}
	return
}

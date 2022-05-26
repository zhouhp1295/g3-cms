package dao

import (
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/auth"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
	"strings"
)

const SuperRoleID = 9999
const SuperRoleIDStr = "9999"

type sysRoleDAO struct {
	crud.BaseDao
}

var SysRoleDao = &sysRoleDAO{
	crud.BaseDao{Model: new(model.SysRole)},
}

func (dao *sysRoleDAO) RefreshRolePerms() {
	setPerms(listRolePermsFromCache())
}

func (dao *sysRoleDAO) GetPermissions(roles string) ([]string, []string) {
	if roles == SuperRoleIDStr {
		return []string{auth.RootUser}, []string{auth.RootPerm}
	}
	resultIdentifiers := make([]string, 0)
	resultPermissions := make([]string, 0)
	roleIds := strings.Split(roles, ",")

	if helpers.IndexOf[string](roleIds, SuperRoleIDStr) >= 0 {
		return []string{auth.RootUser}, []string{auth.RootPerm}
	}

	cachedData := listRolePermsFromCache()
	for _, roleId := range roleIds {
		if data, exist := cachedData[roleId]; exist {
			resultIdentifiers = append(resultIdentifiers, data.Identifier)
			resultPermissions = append(resultPermissions, data.Perms...)
		}
	}
	return resultIdentifiers, resultPermissions
}

func (dao *sysRoleDAO) RoleOptions() []helpers.SelectOption {
	return listRoleOptionsFromCache()
}

func (dao *sysRoleDAO) BeforeInsert(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.SysRole); _ok {
		if dao.CountByColumn("identifier", _m.Identifier) > 0 {
			msg = "标识符已存在"
			return
		}
		ok = true
	}
	return
}

func (dao *sysRoleDAO) BeforeUpdate(m crud.ModelInterface) (ok bool, msg string) {
	if _m, _ok := m.(*model.SysRole); _ok {
		if dao.CountByColumn("identifier", _m.Identifier) > 1 {
			msg = "标识符已存在"
			return
		}
		ok = true
	}
	return
}

func (dao *sysRoleDAO) AfterInsert(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	clearAllRoleCache()
	return
}

func (dao *sysRoleDAO) AfterUpdate(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	clearAllRoleCache()
	return
}
func (dao *sysRoleDAO) AfterRemove(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	clearAllRoleCache()
	return
}
func (dao *sysRoleDAO) AfterDelete(m crud.ModelInterface) (ok bool, msg string) {
	ok = true
	clearAllRoleCache()
	return
}

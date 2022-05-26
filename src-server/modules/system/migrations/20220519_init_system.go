package migrations

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/crud"
	"gorm.io/gorm"
)

var baseSysDictTypeData = `[
	{"name":"用户性别","code":"sys_user_sex","remark":"用户性别列表"},
	{"name":"通用状态","code":"sys_normal_status","remark":"通用状态列表"},
	{"name":"是或否","code":"sys_normal_yes_no","remark":"是或否列表"}
]`

var baseSysDictDataData = `[
	{"label":"男","code":"sys_user_sex","isDefault":"1","value":"1","sort":10,"remark":"男"},
	{"label":"女","code":"sys_user_sex","value":"2","sort":20,"remark":"女"},
	{"label":"未知","code":"sys_user_sex","value":"0","sort":30,"remark":"未知"},
	{"label":"正常","code":"sys_normal_status","isDefault":"1","value":"1","sort":10,"remark":"正常"},
	{"label":"停用","code":"sys_normal_status","value":"0","sort":20,"remark":"停用"},
	{"label":"是","code":"sys_normal_yes_no","isDefault":"1","value":"1","sort":10,"remark":"是"},
	{"label":"否","code":"sys_normal_yes_no","value":"0","sort":20,"remark":"否"}
]`

var baseRoleData = `
[
	{"id":9999, "name":"超级管理员", "identifier":"root", "menus":"", "scope":"9", "sort":0},
	{"id":10000, "name":"管理员", "identifier":"admin", "menus":"", "sort":10},
	{"id":10001, "name":"普通用户", "identifier":"user", "menus":"", "sort":20}
]
`

var baseMenuData = `
[
	{"id":1, "name":"System", "title":"系统管理", "path":"/system", "type":"1", "icon": "system", "component":"Layout", "sort":0},
	{"id":100, "pid":1, "name":"User", "title":"用户管理", "path":"user", "type":"2", "icon": "user", "component":"system/user/index", "perms":"system:user:list", "sort":0},
	{"id":10001, "pid":100, "title":"用户查询", "type":"3", "perms":"system:user:query", "sort":0},
	{"id":10002, "pid":100, "title":"用户新增", "type":"3", "perms":"system:user:add", "sort":1},
	{"id":10003, "pid":100, "title":"用户编辑", "type":"3", "perms":"system:user:edit", "sort":2},
	{"id":10004, "pid":100, "title":"用户删除", "type":"3", "perms":"system:user:remove", "sort":3},
	{"id":101, "pid":1, "name":"Role", "title":"角色管理", "path":"role", "type":"2", "icon": "peoples", "component":"system/role/index", "perms":"system:role:list", "sort":10},
	{"id":10101, "pid":101, "title":"角色查询", "type":"3", "perms":"system:role:query", "sort":0},
	{"id":10102, "pid":101, "title":"角色新增", "type":"3", "perms":"system:role:add", "sort":1},
	{"id":10103, "pid":101, "title":"角色编辑", "type":"3", "perms":"system:role:edit", "sort":2},
	{"id":10104, "pid":101, "title":"角色删除", "type":"3", "perms":"system:role:remove", "sort":3},
	{"id":102, "pid":1, "name":"Config", "title":"参数管理", "path":"config", "type":"2", "icon": "edit", "component":"system/config/index", "perms":"system:config:list", "sort":20},
	{"id":10201, "pid":102, "title":"参数查询", "type":"3", "perms":"system:config:query", "sort":0},
	{"id":10202, "pid":102, "title":"参数新增", "type":"3", "perms":"system:config:add", "sort":1},
	{"id":10203, "pid":102, "title":"参数编辑", "type":"3", "perms":"system:config:edit", "sort":2},
	{"id":10204, "pid":102, "title":"参数删除", "type":"3", "perms":"system:config:remove", "sort":3}
]
`

func CreateSysDictType(db *gorm.DB, data string) error {
	dictTypeList := make([]model.SysDictType, 0)
	err := jsoniter.UnmarshalFromString(data, &dictTypeList)
	if err != nil {
		return errors.Wrap(err, "CreateSysDict UnmarshalFromString Error")
	}
	if len(dictTypeList) > 0 {
		err = db.CreateInBatches(dictTypeList, len(dictTypeList)).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSysDictData(db *gorm.DB, data string) error {
	dictDataList := make([]model.SysDictData, 0)
	err := jsoniter.UnmarshalFromString(data, &dictDataList)
	if err != nil {
		return errors.Wrap(err, "CreateSysDictData UnmarshalFromString Error")
	}
	if len(dictDataList) > 0 {
		err = db.CreateInBatches(dictDataList, len(dictDataList)).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSysRoles(db *gorm.DB, data string) error {
	roleList := make([]model.SysRole, 0)
	err := jsoniter.UnmarshalFromString(data, &roleList)
	if err != nil {
		return errors.Wrap(err, "CreateSysRoles UnmarshalFromString Error")
	}
	if len(roleList) > 0 {
		return db.CreateInBatches(roleList, len(roleList)).Error
	}

	return nil
}

func CreateSystemMenus(db *gorm.DB, data string) error {
	menuList := make([]model.SysMenu, 0)
	err := jsoniter.UnmarshalFromString(data, &menuList)
	if err != nil {
		return errors.Wrap(err, "CreateSystemMenus UnmarshalFromString Error")
	}
	if len(menuList) > 0 {
		return db.CreateInBatches(menuList, len(menuList)).Error
	}

	return nil
}

const M20220519InitSystemCode = "20220519_init_system"

func M20220519InitSystem() func() error {
	return func() error {
		rootDB := crud.DbSess()
		//开启事务
		return rootDB.Transaction(func(tx *gorm.DB) error {
			//初始化字典数据
			err := CreateSysDictType(tx, baseSysDictTypeData)
			if err != nil {
				boot.Logger.Fatal("20220519_init_system , err = %s", err.Error())
				return err
			}
			err = CreateSysDictData(tx, baseSysDictDataData)
			if err != nil {
				boot.Logger.Fatal("20220519_init_system , err = %s", err.Error())
				return err
			}
			//初始化新密码
			config := new(model.SysConfig)
			config.Name = "新用户初始密码"
			config.Code = "sys.user.initPassword"
			config.Value = "123456x"
			config.Remark = "新用户初始密码"
			tx.Create(config)
			//初始化角色
			err = CreateSysRoles(tx, baseRoleData)
			if err != nil {
				boot.Logger.Fatal("20220519_init_system , err = %s", err.Error())
				return err
			}
			err = CreateSystemMenus(tx, baseMenuData)
			if err != nil {
				boot.Logger.Fatal("20220519_init_system , err = %s", err.Error())
				return err
			}
			return nil
		})
	}
}

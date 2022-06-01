package model

import (
	"github.com/zhouhp1295/g3/crud"
)

type SysUser struct {
	crud.BaseModel
	Username string `gorm:"INDEX;TYPE:VARCHAR(20);COMMENT:登录账户" json:"username" form:"username" query:"like"`
	Password string `gorm:"TYPE:VARCHAR(60);COMMENT:密码" json:"password" form:"password"`
	Nickname string `gorm:"TYPE:VARCHAR(20);COMMENT:昵称" json:"nickname" form:"nickname" query:"like"`
	Avatar   string `gorm:"TYPE:VARCHAR(255);COMMENT:头像" json:"avatar" form:"avatar"`
	Phone    string `gorm:"TYPE:VARCHAR(20);COMMENT:联系电话" json:"phone" form:"phone" query:"like"`
	Email    string `gorm:"TYPE:VARCHAR(50);COMMENT:邮箱" json:"email" form:"email"`
	Sex      string `gorm:"TYPE:CHAR(1);NOT NULL;DEFAULT:0;COMMENT:性别" json:"sex" form:"sex" query:"like"`
	IsAdmin  string `gorm:"TYPE:CHAR(1);NOT NULL;DEFAULT:1;COMMENT:是否管理员 0=NO 1=YES"  json:"isAdmin" form:"isAdmin" query:"eq"`
	Roles    string `gorm:"TYPE:VARCHAR(255);COMMENT:用户角色" json:"roles" form:"roles"`
	crud.TailColumns
}

// Table 返回表名
func (*SysUser) Table() string {
	return "sys_user"
}

// NewModel 返回实例
func (*SysUser) NewModel() crud.ModelInterface {
	return new(SysUser)
}

// NewModels 返回实例数组
func (*SysUser) NewModels() interface{} {
	return make([]SysUser, 0)
}

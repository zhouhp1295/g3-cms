package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/common/service"
	"github.com/zhouhp1295/g3-cms/modules/system/dao"
	"github.com/zhouhp1295/g3-cms/modules/system/model"
	"github.com/zhouhp1295/g3/auth"
	"github.com/zhouhp1295/g3/crud"
	"github.com/zhouhp1295/g3/helpers"
	"github.com/zhouhp1295/g3/net"
	"strings"
)

type _sysUserApi struct {
	net.BaseApi
}

var SysUserApi = &_sysUserApi{
	net.BaseApi{Dao: dao.SysUserDao},
}

const (
	PermSystemUserList   = "system:user:list"
	PermSystemUserQuery  = "system:user:query"
	PermSystemUserAdd    = "system:user:add"
	PermSystemUserEdit   = "system:user:edit"
	PermSystemUserRemove = "system:user:remove"
)

func init() {
	boot.SetApiOpen("/admin/system/user/login")
	boot.SetApiOpen("/admin/system/user/logout")

	//基础接口
	boot.ApiGet("/admin/system/user/page", SysUserApi.HandlePage, PermSystemUserQuery)
	boot.ApiGet("/admin/system/user/get", SysUserApi.HandleGet, PermSystemUserQuery)
	boot.ApiPost("/admin/system/user/insert", SysUserApi.HandleInsert, PermSystemUserAdd)
	boot.ApiPut("/admin/system/user/update", SysUserApi.HandleUpdate, PermSystemUserEdit)
	boot.ApiPut("/admin/system/user/status", SysUserApi.HandleUpdateStatus, PermSystemUserEdit)
	boot.ApiDelete("/admin/system/user/delete", SysUserApi.HandleDelete, PermSystemUserRemove)
	//扩展接口
	boot.ApiGet("/admin/system/user/login", SysUserApi.handleLogin)
	boot.ApiPost("/admin/system/user/login", SysUserApi.handleLogin)
	boot.ApiPost("/admin/system/user/logout", SysUserApi.handleLogout)

	boot.ApiGet("/admin/system/user/info", SysUserApi.handleInfo)
	boot.ApiGet("/admin/system/user/routers", SysUserApi.handleRouters)
	//profile
	boot.ApiGet("/admin/system/user/profile/get", SysUserApi.handleGetUserProfile)
	boot.ApiPut("/admin/system/user/profile/update", SysUserApi.handleUpdateUserProfile)
	boot.ApiPut("/admin/system/user/profile/password", SysUserApi.handleUpdateUserPassword)
	boot.ApiPost("/admin/system/user/profile/avatar", SysUserApi.handleUpdateUserAvatar)
}

// userLoginParams 登录参数
type userLoginParams struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Code     string `form:"code" json:"code"`
	UUID     string `form:"uuid" json:"uuid"`
}

// handleLogin 登录
func (api *_sysUserApi) handleLogin(ctx *gin.Context) {
	var params userLoginParams
	var err error
	if err = ctx.ShouldBindJSON(&params); err != nil {
		api.FailedMessage(ctx, "参数错误")
		return
	}
	if !service.CaptchaService.VerifyString(params.UUID, params.Code) {
		api.FailedMessage(ctx, "验证码错误")
		return
	}
	res := api.Dao.FindOneByColumn("username", params.Username)
	if res == nil {
		api.FailedMessage(ctx, "用户不存在")
		return
	}
	user, _ := res.(*model.SysUser)

	if !helpers.PasswordVerify(user.Password, params.Password) {
		api.FailedMessage(ctx, "密码错误")
		return
	}
	if user.Status != crud.FlagYes {
		api.FailedMessage(ctx, "用户已被禁用，请联系管理员")
		return
	}
	var token string
	roleIdentifiers, _ := dao.SysRoleDao.GetPermissions(user.Roles)
	if token, err = boot.GinCtx().ApiRouter.JwtAuth.Token(user.Id, strings.Join(roleIdentifiers, ",")); err != nil {
		api.FailedServerError(ctx, err.Error(), "")
		return
	}
	api.SuccessData(ctx, gin.H{"token": token})
}

// handleLogout 注销
func (api *_sysUserApi) handleLogout(ctx *gin.Context) {
	api.SuccessDefault(ctx)
}

// handleInfo 用户信息
func (api *_sysUserApi) handleInfo(ctx *gin.Context) {
	uid := ctx.GetInt64(auth.CtxJwtUid)

	res := api.Dao.FindByPk(uid)
	if res == nil {
		api.FailedMessage(ctx, "用户不存在")
		return
	}
	user, _ := res.(*model.SysUser)

	result := gin.H{
		"user": user,
	}
	result["roles"], result["permissions"] = dao.SysRoleDao.GetPermissions(user.Roles)
	api.SuccessData(ctx, result)
}

// handleRouters 获取用户的可用vue可用路由
func (api *_sysUserApi) handleRouters(ctx *gin.Context) {
	uid := ctx.GetInt64(auth.CtxJwtUid)
	res := api.Dao.FindByPk(uid)
	if res == nil {
		api.FailedMessage(ctx, "用户不存在")
		return
	}
	user, _ := res.(*model.SysUser)

	_, permissionList := dao.SysRoleDao.GetPermissions(user.Roles)
	if len(permissionList) == 1 && permissionList[0] == auth.RootPerm {
		api.SuccessData(ctx, gin.H{"routers": dao.SysMenuDao.GetMenuRouters()})
		return
	}
	api.SuccessData(ctx, gin.H{"routers": dao.SysMenuDao.GetMenuRoutersByPerms(permissionList)})
}

// handleGetUserProfile 获取个人主页信息
func (api *_sysUserApi) handleGetUserProfile(ctx *gin.Context) {
	uid := ctx.GetInt64(auth.CtxJwtUid)
	api.SuccessData(ctx, dao.SysUserDao.FindByPk(uid))
}

type updateUserProfileParams struct {
	Username string `json:"username" form:"username"`
	Sex      string `json:"sex" form:"sex"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
}

// handleUpdateUserProfile 更新个人主页信息
func (api *_sysUserApi) handleUpdateUserProfile(ctx *gin.Context) {
	uid := ctx.GetInt64(auth.CtxJwtUid)
	params := updateUserProfileParams{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		api.FailedMessage(ctx, "参数错误:"+err.Error())
		return
	}
	user := new(model.SysUser)
	user.Id = uid
	user.Username = params.Username
	user.Sex = params.Sex
	user.Phone = params.Phone
	user.Email = params.Email
	if api.Dao.Update(user, uid) {
		api.SuccessDefault(ctx)
	} else {
		api.FailedMessage(ctx, "操作失败,请稍后重试")
	}
}

type updateUserPasswordParams struct {
	OldPassword string `json:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" form:"newPassword"`
}

// handleUpdatePassword 更新个人密码
func (api *_sysUserApi) handleUpdateUserPassword(ctx *gin.Context) {
	uid := ctx.GetInt64(auth.CtxJwtUid)
	params := updateUserPasswordParams{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		api.FailedMessage(ctx, "参数错误:"+err.Error())
		return
	}
	res := api.Dao.FindByPk(uid)
	if res == nil {
		api.FailedMessage(ctx, "用户不存在")
		return
	}
	user, _ := res.(*model.SysUser)

	if !helpers.PasswordVerify(user.Password, params.OldPassword) {
		api.FailedMessage(ctx, "旧密码错误")
		return
	}
	hashPwd, err := helpers.PasswordHash(params.NewPassword)
	if err != nil {
		api.FailedMessage(ctx, "密码加密错误")
		return
	}
	if api.Dao.UpdateColumn(uid, "password", hashPwd, uid) {
		api.SuccessDefault(ctx)
	} else {
		api.FailedMessage(ctx, "操作失败,请稍后重试")
	}
}

func (api *_sysUserApi) handleUpdateUserAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("avatar")
	if err != nil {
		boot.Logger.Error("handleUpdateUserAvatar err = %s\n", err.Error())
		api.FailedMessage(ctx, "操作失败:"+err.Error())
		return
	}

	filePath, msg, ok := service.UploadService.UploadImage(file)
	if !ok {
		boot.Logger.Error("handleUpdateUserAvatar err = %s\n", msg)
		api.FailedMessage(ctx, "操作失败:"+msg)
		return
	}
	uid := ctx.GetInt64(auth.CtxJwtUid)

	user := new(model.SysUser)
	user.Id = uid
	user.Avatar = filePath
	if api.Dao.Update(user, uid) {
		api.SuccessData(ctx, gin.H{"avatar": filePath})
	} else {
		api.FailedMessage(ctx, "操作失败,请稍后重试")
	}
}

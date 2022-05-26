package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3/auth"
	"github.com/zhouhp1295/g3/net"
	"gopkg.in/ini.v1"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
)

var File *ini.File

type GinRouter struct {
	Group   *gin.RouterGroup
	Perms   *auth.Perm
	JwtAuth *auth.JwtAuth
}

func (ginRouter *GinRouter) Use(middleware ...gin.HandlerFunc) {
	ginRouter.Group.Use(middleware...)
}

func (ginRouter *GinRouter) AddOpenRouter(routers ...string) {
	ginRouter.JwtAuth.AddOpenRouters(routers...)
}

func (ginRouter *GinRouter) AddWhiteRouter(routers ...string) {
	ginRouter.JwtAuth.AddWhiteRouters(routers...)
}

func (ginRouter *GinRouter) AddRouter(method, router string, handler gin.HandlerFunc, perms ...string) {
	ginRouter.Group.Handle(method, router, handler)
	if len(perms) > 0 {
		ginRouter.Perms.AddRouterPerms(router, perms...)
	}
}

func newApiRouter(r *gin.Engine) *GinRouter {
	apiRouter := &GinRouter{
		Group: r.Group("/api"),
		Perms: auth.NewPerm(),
	}
	apiRouter.JwtAuth = auth.NewJwt("/api", apiRouter.Perms, JwtCfg.Secret, JwtCfg.ExpiredSeconds)

	apiRouter.Use(apiRouter.JwtAuth.Authentication)

	return apiRouter
}

type ginContext struct {
	Engine    *gin.Engine
	ApiRouter *GinRouter
}

var ginCtx *ginContext
var ginCtxOnce sync.Once

func BindWebStatic() {
	GinCtx().Engine.Static("/backend", path.Join(HomeDir(), "public", "backend"))
	GinCtx().Engine.Static("/upload", path.Join(HomeDir(), "upload"))
	GinCtx().Engine.StaticFile("/favicon.ico", path.Join(HomeDir(), "public", "favicon.ico"))
}

func GinCtx() *ginContext {
	ginCtxOnce.Do(func() {
		gin.DefaultWriter = io.MultiWriter(ginLoggerWriter, os.Stdout)

		r := gin.Default()
		r.Use(net.DefaultCors())

		ginCtx = &ginContext{
			Engine:    r,
			ApiRouter: newApiRouter(r),
		}
	})

	return ginCtx
}

func ApiGet(router string, handler gin.HandlerFunc, perms ...string) {
	GinCtx().ApiRouter.AddRouter(http.MethodGet, router, handler, perms...)
	if len(perms) == 0 {
		SetApiWhite(router)
	}
}
func ApiPut(router string, handler gin.HandlerFunc, perms ...string) {
	GinCtx().ApiRouter.AddRouter(http.MethodPut, router, handler, perms...)
	if len(perms) == 0 {
		SetApiWhite(router)
	}
}
func ApiPost(router string, handler gin.HandlerFunc, perms ...string) {
	GinCtx().ApiRouter.AddRouter(http.MethodPost, router, handler, perms...)
	if len(perms) == 0 {
		SetApiWhite(router)
	}
}
func ApiDelete(router string, handler gin.HandlerFunc, perms ...string) {
	GinCtx().ApiRouter.AddRouter(http.MethodDelete, router, handler, perms...)
	if len(perms) == 0 {
		SetApiWhite(router)
	}
}

func SetApiWhite(routers ...string) {
	GinCtx().ApiRouter.AddWhiteRouter(routers...)
}

func SetApiOpen(routers ...string) {
	GinCtx().ApiRouter.AddOpenRouter(routers...)
}

// IsProdMode returns true if the application is running in production mode.
func IsProdMode() bool {
	return strings.EqualFold(App.RunMode, "prod")
}

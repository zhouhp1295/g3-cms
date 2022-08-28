// Copyright (c) 554949297@qq.com . 2022-2022. All rights reserved

//go:build http
// +build http

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/common/service"
	"github.com/zhouhp1295/g3/net"
	"net/http"
)

func init() {
	boot.RegisterAfterInstallFunction(func() {
		g3.GetGin().Group("/api").MakeOpen("/common/captcha")
		g3.GetGin().Group("/api").
			Bind(http.MethodGet, "/common/captcha", handleCaptcha)
	})
}

func handleCaptcha(ctx *gin.Context) {
	captchaId, captchaBase64 := service.CaptchaService.GetStringImg()
	net.SuccessData(ctx, gin.H{
		"id":     captchaId,
		"base64": captchaBase64,
	})
}

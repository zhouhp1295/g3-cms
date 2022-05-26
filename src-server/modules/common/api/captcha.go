package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouhp1295/g3-cms/boot"
	"github.com/zhouhp1295/g3-cms/modules/common/service"
)

func init() {
	boot.SetApiOpen("/common/captcha")
	boot.ApiGet("/common/captcha", handleCaptcha)
}

func handleCaptcha(ctx *gin.Context) {
	captchaId, captchaBase64 := service.CaptchaService.GetStringImg()
	commonApi.SuccessData(ctx, gin.H{
		"id":     captchaId,
		"base64": captchaBase64,
	})
}

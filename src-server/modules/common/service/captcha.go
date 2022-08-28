// Copyright (c) 554949297@qq.com . 2022-2022 . All rights reserved

package service

import (
	"github.com/mojocn/base64Captcha"
	"github.com/zhouhp1295/g3"
	"go.uber.org/zap"
	"strings"
)

type captchaService struct {
}

var CaptchaService = new(captchaService)

func (service *captchaService) GetStringImg() (captchaId, captchaBase64 string) {
	driver := &base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      50,
		ShowLineOptions: 20,
		Length:          4,
		Source:          "abcdefghjkmnpqrstuvwxyz",
		Fonts:           []string{"chromohv.ttf"},
	}
	driver = driver.ConvertFonts()
	store := base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)
	captchaId, captchaBase64, err := c.Generate()
	if err != nil {
		g3.ZL().Error("GetStringImg", zap.Error(err))
	}
	return
}

//VerifyString 验证输入的验证码是否正确
func (service *captchaService) VerifyString(id, answer string) bool {
	driver := new(base64Captcha.DriverString)
	store := base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)
	answer = strings.ToLower(answer)
	return c.Verify(id, answer, true)
}

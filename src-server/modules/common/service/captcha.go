package service

import (
	"github.com/mojocn/base64Captcha"
	"github.com/zhouhp1295/g3-cms/boot"
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
		boot.Logger.Error("GetStringImg err = %s\n", err.Error())
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

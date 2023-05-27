package controller

import (
	"image/color"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"

	"project/utils"
)

func GenerateCaptcha(c *gin.Context) {
	driver := &base64Captcha.DriverString{
		Height:          36,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 36,
			G: 136,
			B: 188,
			A: 159,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	captcha := base64Captcha.NewCaptcha(driver, utils.Redis_Store)
	id, b64, err := captcha.Generate()
	if err != nil {
		zap.S().Error("验证码获取失败: ", err.Error())
		return
	}
	utils.Success(c, gin.H{
		"id":   id,
		"path": b64,
	}, "验证码已发送")
}

func VerifyCaptcha(id, value string) bool {
	return utils.Redis_Store.Verify(id, value, true)
}

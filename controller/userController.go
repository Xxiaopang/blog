package controller

import (
	"net/http"
	"project/vo"

	"github.com/gin-gonic/gin"

	"project/config"
	"project/service"
	"project/utils"
)

func Register(c *gin.Context) {
	register := vo.Register{}

	if err := c.ShouldBindJSON(&register); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "ShouldBindJSON 解析失败"+err.Error())
		return
	}
	if len(register.Mobile) != 11 {
		utils.Response(c, http.StatusUnprocessableEntity, nil, "手机号必须为11位")
		return
	}
	if len(register.Password) < 6 {
		utils.Response(c, http.StatusUnprocessableEntity, nil, "密码不能少于6位")
		return
	}
	if len(register.Name) == 0 {
		register.Name = "用户" + register.Mobile
	}

	if _, ok := service.FindUserByMobile(register.Mobile); ok {
		utils.Fail(c, nil, "用户已经存在")
		return
	}
	hash, err := utils.Encryption(register.Password)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "加密错误")
		return
	}
	register.Password = string(hash)
	err = service.CreateUser(register.ToUser())
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	utils.Success(c, nil, "注册成功")
}

func Login(c *gin.Context) {
	var login vo.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "ShouldBindJSON 解析失败"+err.Error())
		return
	}

	if !VerifyCaptcha(login.CaptchaID, login.Captcha) {
		utils.Response(c, http.StatusUnprocessableEntity, nil, "验证码错误")
		return
	}
	if len(login.Mobile) != 11 {
		utils.Response(c, http.StatusUnprocessableEntity, nil, "手机号必须为11位")
		return
	}
	user, ok := service.FindUserByMobile(login.Mobile)
	if !ok {
		utils.Fail(c, nil, "用户尚未注册")
		return
	}

	if err := utils.Decrypt([]byte(user.Password), []byte(login.Password)); err != nil {
		utils.Fail(c, nil, "密码错误")
		return
	}
	token, err := utils.CreateToken(user.UID, user.RoleId)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "系统异常")
		return
	}
	c.Header("Authorization", config.Header+"$"+token)

	utils.Success(c, token, "登录成功")
}

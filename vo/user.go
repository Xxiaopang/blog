package vo

import "project/model"

type Register struct {
	Name     string `json:"username" `
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

func (this *Register) ToUser() *model.User {
	return &model.User{
		Name:     this.Name,
		Password: this.Password,
		Mobile:   this.Mobile,
		Email:    this.Email,
	}
}

type Login struct {
	Mobile    string `json:"mobile,omitempty"`
	Password  string `json:"password,omitempty"`
	CaptchaID string `json:"captcha_id,omitempty"`
	Captcha   string `json:"captcha,omitempty"`
}

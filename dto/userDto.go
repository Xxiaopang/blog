package dto

import (
	"project/model"
)

type User struct {
	ID       uint
	Name     string
	Password string
	Mobile   string
	Email    string
	RoleId   int
	Status   int
}

func (this *User) ToUser() *model.User {
	return &model.User{
		Name:     this.Name,
		Password: this.Password,
		Mobile:   this.Mobile,
		Email:    this.Email,
	}
}

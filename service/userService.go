package service

import (
	"github.com/google/uuid"
	"project/database"
	"project/model"
)

func FindUserByID(id uint) (*model.User, error) {
	user := model.User{}
	err := database.DB.Where("id = ?", id).First(&user).Error
	return &user, err

}

func FindUserByMobile(mobile string) (*model.User, bool) {
	u := model.User{}
	database.DB.Where("mobile = ?", mobile).First(&u)
	if u.ID != 0 {
		return &u, true
	}
	return nil, false
}

func FindAllUser() (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func CreateUser(u *model.User) error {
	u.UID = uint(uuid.New().ID())
	return database.DB.Create(&u).Error
}

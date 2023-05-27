package service

import (
	"github.com/google/uuid"
	"project/database"
	"project/model"
)

func FindAllTags(page, limit int) (tags []*model.Tag, num int64) {
	num = database.DB.Find(&tags).Offset((page - 1) * limit).Limit(limit).RowsAffected
	return
}

func FindTags(name string) (tag model.Tag, err error) {
	err = database.DB.Where("name = ?", name).Find(&tag).Error
	return
}

func FindTagByID(uuid uint) (tag model.Tag, err error) {
	err = database.DB.Where("uuid = ?", uuid).Find(&tag).Error
	return
}

func CreateTags(name string) error {
	return database.DB.Create(&model.Tag{
		Name: name,
		UUID: uint(uuid.New().ID()),
	}).Error
}

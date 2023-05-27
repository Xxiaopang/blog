package service

import (
	"github.com/google/uuid"
	"project/database"
	"project/dto"
	"project/model"
)

func CreateFile(oldName, newName string) error {
	return database.DB.Create(&model.File{
		UUID:    uint(uuid.New().ID()),
		OldName: oldName,
		NewName: newName,
	}).Error
}

func FindFileByID(id uint) (dto.FileDto, error) {
	file := model.File{}
	err := database.DB.Where("uuid = ?", id).Find(&file).Error
	fileDto := dto.ToFileDto(file)
	return fileDto, err
}

func FindAllFile(page, limit int) (*[]dto.FileDto, int64) {
	var list []model.File
	rowsAffected := database.DB.Find(&list).Offset((page - 1) * limit).Limit(limit).RowsAffected
	fileDtos := dto.ToManyFileDto(list)
	return fileDtos, rowsAffected
}

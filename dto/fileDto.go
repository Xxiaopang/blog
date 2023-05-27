package dto

import "project/model"

type FileDto struct {
	UUID    uint   `json:"uuid,omitempty"`
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
}

func ToFileDto(file model.File) FileDto {
	return FileDto{
		UUID:    file.UUID,
		OldName: file.OldName,
		NewName: file.NewName,
	}
}
func toFileDto(file model.File) FileDto {
	return FileDto{
		UUID:    file.UUID,
		OldName: file.OldName,
	}
}

func ToManyFileDto(files []model.File) *[]FileDto {
	var fileDtos = make([]FileDto, len(files))
	for i := 0; i < len(files); i++ {
		fileDtos[i] = toFileDto(files[i])
	}
	return &fileDtos
}

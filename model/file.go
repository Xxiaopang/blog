package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
	UUID    uint   `json:"uuid,omitempty" gorm:"not null;unique"`
}

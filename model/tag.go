package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `json:"name,omitempty;unique"`
	UUID uint   `json:"uuid,omitempty" gorm:"not null;unique"`
	//Articles []Article `gorm:"many2many:article_tag"`
	Articles []Article `json:"articles" gorm:"many2many:article_tag"`
}

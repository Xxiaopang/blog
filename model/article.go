package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UUID     uint      `json:"uuid,omitempty" gorm:"not null;;unique;comment:'文章uuid'"`
	UID      uint      `json:"uid,omitempty" gorm:"not null;comment:'用户'"`
	Title    string    `json:"title,omitempty" gorm:"type:text;not null"`
	Content  string    `json:"content,omitempty" gorm:"type:text;not null"`
	Likes    int       `json:"likes"`
	Comments []Comment `json:"comments" gorm:"many2many:Article_comments;"`
	Tags     []Tag     `json:"tags" gorm:"many2many:article_tag"`
}

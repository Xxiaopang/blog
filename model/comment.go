package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string `json:"content"`    // 评论内容
	ArticleID uint   `json:"article_id"` // 所属文章ID
}

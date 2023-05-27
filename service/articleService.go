package service

import (
	"github.com/google/uuid"
	"project/database"
	"project/model"
)

// - GET /articles:获取所有文章列表
// - GET /articles/:id:获取指定文章
// - POST /articles:创建新文章
// - PUT /articles/:id:更新文章
// - DELETE /articles/:id:删除文章

//- GET /articles/:id/comments:获取指定文章的评论
//- POST /articles/:id/comments:为指定文章创建评论

func FindAllArticle(page, limit int) ([]*model.Article, error) {
	var article []*model.Article
	err := database.DB.Find(&article).Offset((page - 1) * limit).Limit(limit).Error
	return article, err
}

func FindArticleByTag(tag model.Tag, page, limit int) ([]model.Article, int64, error) {
	var articles []model.Article
	var count int64
	//rowsAffected := database.DB.Preload("tags").Find(&article).Offset((page - 1) * limit).Limit(limit).RowsAffected
	//err := database.DB.Model(&tag).Association("Articles").Find(&articles)
	err := database.DB.Model(&tag).Offset((page - 1) * limit).Limit(limit).Association("Articles").Find(&articles)
	if err != nil {
		return nil, 0, err
	}
	count = database.DB.Model(&tag).Association("Articles").Count()
	return articles, count, err
}

func FindArticleByUser(uid uint, page, limit int) ([]model.Article, int64, error) {
	var articles []model.Article
	var count int64

	err := database.DB.Where("uid = ?", uid).Offset((page - 1) * limit).Limit(limit).Find(&articles).Count(&count).Error

	return articles, count, err
}

func FindArticleByID(id uint) (model.Article, error) {
	article := model.Article{}
	err := database.DB.Preload("Tags").Where("uuid = ?", id).Find(&article).Error
	return article, err
}

func CreateArticle(article model.Article) (err error) {
	article.UUID = uint(uuid.New().ID())
	tx := database.DB.Begin()

	err = tx.Create(&article).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit().Error; err != nil {
		return
	}
	return
}

func UpdateArticle(article model.Article) (newArticle model.Article, err error) {

	tx := database.DB.Begin()
	oldArticle, err := FindArticleByID(article.UUID)
	if err != nil {
		return
	}
	if len(oldArticle.Tags) != 0 {

		err = tx.Model(&oldArticle).Association("Tags").Delete(oldArticle.Tags)
		if err != nil {
			return
		}
	}
	if len(article.Tags) != 0 {
		err = tx.Model(&oldArticle).Association("Tags").Append(article.Tags)
		if err != nil {
			return
		}
	}
	err = tx.Model(oldArticle).Updates(map[string]interface{}{
		"title":   article.Title,
		"Content": article.Content,
	}).Error
	if err = tx.Commit().Error; err != nil {
		return
	}
	return FindArticleByID(article.UUID)

}

func DeleteArticle(id uint) error {
	article, err := FindArticleByID(id)
	if err != nil {
		return err
	}
	if len(article.Tags) != 0 {
		err = database.DB.Model(&article).Association("Tags").Delete(article.Tags)
		if err != nil {
			return err
		}
	}

	return database.DB.Where("uuid = ?", id).Delete(&model.Article{}).Error
}

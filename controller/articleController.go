package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"project/model"
	"project/service"
	"project/utils"
	"project/vo"
)

// article.GET("/:id", controller.FindArticleByID)
func FindArticleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		zap.S().Info("string to int fail" + err.Error())
		return
	}
	article, err := service.FindArticleByID(uint(id))
	if err != nil {
		zap.S().Info("service.FindArticleByID" + err.Error())
		utils.Success(c, nil, "查询失败")
		return
	}
	utils.Success(c, article, "查询成功")
}

func FindAllArticle(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		zap.S().Info("string to int fail" + err.Error())
		utils.Success(c, nil, "转换失败")
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		zap.S().Info("string to int fail" + err.Error())
		utils.Success(c, nil, "转换失败")
		return
	}
	articles, err := service.FindAllArticle(limit, page)
	if err != nil {
		zap.S().Info("查询失败" + err.Error())
		utils.Success(c, nil, "查询失败")
		return
	}
	utils.Success(c, articles, len(articles))
}

func FindArticleByTag(c *gin.Context) {
	TagStr := c.Query("tag")
	tag, err := service.FindTags(TagStr)
	if err != nil {
		zap.S().Info("标签查询失败 " + err.Error())
		utils.Fail(c, nil, "标签查询失败")
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		zap.S().Info("string to int fail " + err.Error())
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		zap.S().Info("string to int fail " + err.Error())
		utils.Success(c, nil, "查询失败")
		return
	}

	articles, count, err := service.FindArticleByTag(tag, page, limit)
	if err != nil {
		zap.S().Info("查询失败" + err.Error())
		utils.Fail(c, nil, "查询失败")
		return
	}
	utils.Success(c, articles, count)
}

func FindArticleByUser(c *gin.Context) {
	idStr := c.Param("id")
	uid, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		zap.S().Info("string to int fail" + err.Error())
		return
	}
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		zap.S().Info("string to int fail " + err.Error())
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		zap.S().Info("string to int fail " + err.Error())
		utils.Success(c, nil, "查询失败")
		return
	}

	articles, count, err := service.FindArticleByUser(uint(uid), page, limit)
	if err != nil {
		zap.S().Info("查询失败 " + err.Error())
		utils.Success(c, nil, "查询失败")
		return
	}
	utils.Success(c, articles, count)
}

// article.POST("/", controller.CreateArticle)
func CreateArticle(c *gin.Context) {
	var articleReq vo.ArticleRequest
	if err := c.ShouldBindJSON(&articleReq); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "ShouldBindJSON 解析失败"+err.Error())
		zap.S().Info("ShouldBindJSON 解析失败" + err.Error())
		return
	}
	uid, _ := c.Get("uid")

	article := model.Article{
		UID:     uid.(uint),
		Title:   articleReq.Title,
		Content: articleReq.Content,
	}

	for _, tag := range articleReq.Tags {
		findTags, err := service.FindTags(tag.Name)
		if err != nil {
			return
		}
		article.Tags = append(article.Tags, findTags)

	}

	err := service.CreateArticle(article)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	utils.Success(c, gin.H{"article": article}, "文章发布成功")
}

// article.PUT("/:id", controller.UpdateArticle)
func UpdateArticle(c *gin.Context) {
	var articleReq vo.ArticleRequest
	if err := c.ShouldBindJSON(&articleReq); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "ShouldBindJSON 解析失败"+err.Error())
		zap.S().Info("ShouldBindJSON 解析失败" + err.Error())
		return
	}

	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		zap.S().Error("string to int fail" + err.Error())
		return
	}
	article, err := service.FindArticleByID(uint(id))
	if err != nil {
		utils.Success(c, nil, "文章不存在")
		return
	}
	article.Title = articleReq.Title
	article.Content = articleReq.Content

	article.Tags = make([]model.Tag, 0)
	for _, tag := range articleReq.Tags {
		findTags, err := service.FindTags(tag.Name)
		if err != nil {
			return
		}
		article.Tags = append(article.Tags, findTags)
	}

	newArticle, err := service.UpdateArticle(article)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, err.Error())
		return
	}
	utils.Success(c, gin.H{"article": newArticle}, "文章修改成功")
}

// article.DELETE("/:id", controller.DeleteArticle)
func DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		zap.S().Info("string to int fail" + err.Error())
		return
	}

	err = service.DeleteArticle(uint(id))
	if err != nil {
		utils.Fail(c, "删除失败请重试", nil)
		return
	}

	utils.Success(c, nil, "文章删除成功")
}

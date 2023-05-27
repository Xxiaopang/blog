package router

import (
	"github.com/gin-gonic/gin"
	"project/controller"

	"project/middleware"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Log(), gin.Recovery())

	r.MaxMultipartMemory = 64 << 20

	r.POST("api/login", controller.Login)
	r.POST("api/register", controller.Register)
	r.GET("api/captcha", controller.GenerateCaptcha)

	files := r.Group("api/file")
	{
		files.POST("/upload", controller.UploadFile)
		files.GET("/download/:id", controller.DownloadFile)
		files.GET("/list", controller.FindAllFile)
	}

	r.Use(middleware.JWTAuto(), middleware.Cors())

	r.GET("api/home", func(c *gin.Context) {
		c.JSON(200, "llll")
	})

	//- GET /articles/:id/comments:获取指定文章的评论
	//- POST /articles/:id/comments:为指定文章创建评论
	article := r.Group("api/articles")
	{
		article.GET("/", controller.FindAllArticle)
		article.GET("/:id", controller.FindArticleByID)
		article.GET("/user/:id", controller.FindArticleByUser)
		//article.GET("/tag/:tag", controller.FindArticleByTag)
		article.POST("/", controller.CreateArticle)
		article.PUT("/:id", controller.UpdateArticle)
		article.DELETE("/:id", controller.DeleteArticle)
	}
	r.GET("api/article/", controller.FindArticleByTag)
	tags := r.Group("api/tags")
	{
		tags.GET("/", controller.FindAllTags)
		tags.POST("/", controller.CreateTags)
		//tags.GET("/:id", controller.FindTagsByID)
	}

	return r
}

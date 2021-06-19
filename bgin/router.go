package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/setting"
)

// SetUpRoot ...
func SetUpRoot(root *gin.Engine, env *setting.Environment) {
	root.Use(MiddlewareLogger())
	root.Use(MiddlewareUsecase(env))

	{
		gArticles := root.Group("articles")
		gArticles.GET("", HandlerGetArticles())
		{
			gArticle := gArticles.Group(":articleID")
			gArticle.Use(MiddlewareGetArticle())
			gArticle.GET("", HandlerGetArticlesByID())
		}
	}

	{
		gTags := root.Group("tags")
		gTags.GET("", HandlerGetTags())
		{
			gTag := gTags.Group(":tagID")
			gTag.Use(MiddlewareGetTag())
			gTag.GET("", HandlerGetTagsByID())
			gTag.GET("articles", HandlerGetArticles())
		}
	}

	{
		gAdmin := root.Group("admin")
		gAdmin.Use(MiddlewareAdminAuth())
		gAdmin.GET("sitemap.xml", HandlerGetSitemapXML())
	}
}

package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/application"
	"github.com/suzuito/blog1-go/setting"
)

// SetUpRoot ...
func SetUpRoot(root *gin.Engine, env *setting.Environment, app *application.Application) {
	root.Use(MiddlewareLogger(app))
	root.Use(MiddlewareUsecase(app, env))

	{
		gArticles := root.Group("articles")
		gArticles.GET("", HandlerGetArticles(app))
		{
			gArticle := gArticles.Group(":articleID")
			gArticle.Use(MiddlewareGetArticle(app))
			gArticle.GET("", HandlerGetArticlesByID(app))
		}
	}

	{
		gTags := root.Group("tags")
		gTags.GET("", HandlerGetTags(app))
		{
			gTag := gTags.Group(":tagID")
			gTag.Use(MiddlewareGetTag(app))
			gTag.GET("", HandlerGetTagsByID(app))
			gTag.GET("articles", HandlerGetArticles(app))
		}
	}

	{
		gAdmin := root.Group("admin")
		gAdmin.Use(MiddlewareAdminAuth(app))
		gAdmin.GET("sitemap.xml", HandlerGetSitemapXML(app))
	}
}

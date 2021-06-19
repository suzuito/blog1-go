package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/setting"
)

// SetUpRoot ...
func SetUpRoot(root *gin.Engine, env *setting.Environment) {
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
		gAdmin := root.Group("admin")
		gAdmin.Use(MiddlewareAdminAuth())
		gAdmin.GET("sitemap.xml", HandlerGetSitemapXML())
	}
}

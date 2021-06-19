package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/inject"
	"github.com/suzuito/blog1-go/setting"
)

// SetUpRoot ...
func SetUpRoot(root *gin.Engine, env *setting.Environment, gdeps *inject.GlobalDepends) {
	root.Use(MiddlewareUsecase(env, gdeps))

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

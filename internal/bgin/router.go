package bgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/internal/inject"
	"github.com/suzuito/blog1-go/internal/setting"
)

// SetUpRoot ...
func SetUpRoot(root *gin.Engine, env *setting.Environment, gdeps *inject.GlobalDepends) {
	root.GET("", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	root.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
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

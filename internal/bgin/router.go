package bgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/setting"
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

	root.Static("css", fmt.Sprintf("%s", env.DirPathCSS))
	{
		root.LoadHTMLGlob(fmt.Sprintf("%s/*.html", env.DirPathTemplate))
		gArticles := root.Group("articles")
		gArticles.GET("", HTMLGetArticles())
		{
			gArticle := gArticles.Group(":articleID")
			gArticle.Use(MiddlewareGetArticle())
			gArticle.GET("", HTMLGetArticle())
			gArticle.PUT("", HandlerPutArticleByID())
		}
	}

	{
		gAdmin := root.Group("admin")
		gAdmin.Use(MiddlewareAdminAuth())
		gAdmin.GET("sitemap.xml", HandlerGetSitemapXML())
	}
}

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
	root.Static("css", fmt.Sprintf("%s", env.DirPathCSS))

	root.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	root.Use(MiddlewareUsecase(env, gdeps))

	root.Static("asset", env.DirPathAsset)
	root.GET("sitemap.xml", HandlerGetSitemapXML(env))
	root.GET("robots.txt", HandlerHTMLGetRobots(env))
	root.GET("", HandlerHTMLGetTop(env))
	root.GET("about", HandlerHTMLGetAbout(env))
	root.GET("sandbox", HandlerHTMLGetSandbox(env))
	{
		root.LoadHTMLGlob(fmt.Sprintf("%s/*.html", env.DirPathTemplate))
		gArticles := root.Group("articles")
		gArticles.GET("", HandlerHTMLGetArticles(env))
		{
			gArticle := gArticles.Group(":articleID")
			gArticle.Use(HTMLMiddlewareGetArticle(env))
			gArticle.GET("", HandlerHTMLGetArticle(env))
		}
	}
}

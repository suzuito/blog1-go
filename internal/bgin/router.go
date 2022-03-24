package bgin

import (
	"fmt"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/setting"
)

// SetUpRoot ...
func SetUpRoot(root *gin.Engine, gdeps *inject.GlobalDepends) {
	root.Use(sentrygin.New(sentrygin.Options{}))
	root.Static("css", fmt.Sprintf("%s", setting.E.DirPathCSS))

	root.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	root.Use(MiddlewareUsecase(gdeps))

	root.Static("asset", setting.E.DirPathAsset)
	root.GET("sitemap.xml", HandlerGetSitemapXML())
	root.GET("robots.txt", HandlerHTMLGetRobots())
	root.GET("", HandlerHTMLGetTop())
	root.GET("about", HandlerHTMLGetAbout())
	root.GET("sandbox", HandlerHTMLGetSandbox())
	{
		root.LoadHTMLGlob(fmt.Sprintf("%s/*.html", setting.E.DirPathTemplate))
		gArticles := root.Group("articles")
		gArticles.GET("", HandlerHTMLGetArticles())
		{
			gArticle := gArticles.Group(":articleID")
			gArticle.Use(HTMLMiddlewareGetArticle())
			gArticle.GET("", HandlerHTMLGetArticle())
		}
	}
}

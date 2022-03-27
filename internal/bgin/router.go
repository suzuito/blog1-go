package bgin

import (
	"fmt"
	"html/template"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

var tmplArticle *template.Template

func init() {
	var err error
	tmplArticle, _ = template.New("hoge").ParseGlob(fmt.Sprintf("%s/*.html", setting.E.DirPathTemplate))
	if err != nil {
		panic(errors.Wrapf(err, "cannot new template"))
	}
}

// SetUpRoot ...
func SetUpRoot(root *gin.Engine, u usecase.Usecase) {
	root.Use(sentrygin.New(sentrygin.Options{}))
	root.Static("css", fmt.Sprintf("%s", setting.E.DirPathCSS))

	root.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	root.Use(MiddlewareUsecase(u))

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

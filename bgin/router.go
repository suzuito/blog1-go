package bgin

import "github.com/gin-gonic/gin"

// SetUpRoot ...
func SetUpRoot(root *gin.Engine) {
	root.Use(MiddlewareLogger())

	gArticles := root.Group("articles")
	gArticles.GET(
		"",
		HandlerGetArticles(),
	)
	gArticle := gArticles.Group(":articleID")
	gArticle.Use(MiddlewareGetArticle())
	gArticle.GET(
		"",
		HandlerGetArticlesByID(),
	)
	gTags := root.Group("tags")
	gTags.GET(
		"",
		HandlerGetTags(),
	)
	gTag := gTags.Group(":tagID")
	gTag.Use(MiddlewareGetTag())
	gTag.GET(
		"",
		HandlerGetTagsByID(),
	)
	gTag.GET(
		"articles",
		HandlerGetArticles(),
	)
}

package bgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTMLGetArticle ...
func HTMLGetArticle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		article := getCtxArticle(ctx)
		u := getCtxUsecase(ctx)
		content := []byte{}
		if err := u.GetArticleHTML(ctx, article.ID, &content); err != nil {
			fmt.Println(err)
			ctx.HTML(
				http.StatusNotFound,
				"pc_404.html",
				gin.H{},
			)
			return
		}
		ctx.HTML(
			http.StatusOK,
			"pc_article.html",
			gin.H{
				"Global":  htmlGlobal,
				"Article": article,
				"Content": string(content),
			},
		)
	}
}

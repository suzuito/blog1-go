package bgin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/entity/model"
	env "github.com/suzuito/common-env"
)

// HandlerGetArticles ...
func HandlerGetArticles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := GetCtxUsecase(ctx)
		n := env.GetenvAsInt("n", 10)
		startPublishedAt := env.GetenvAsInt64("start", -1)
		if startPublishedAt < 0 {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				"Param 'start' must be set",
			)
			return
		}
		articles := []model.Article{}
		if err := u.GetArticles(ctx, startPublishedAt, n, &articles); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				err.Error(),
			)
			return
		}
		ctx.JSON(
			http.StatusOK,
			NewResponseArticles(&articles),
		)
	}
}

// HandlerGetArticlesByID ...
func HandlerGetArticlesByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		article := GetCtxArticle(ctx)
		ctx.JSON(
			http.StatusOK,
			NewResponseArticle(article),
		)
	}
}

// HandlerPostArticles ...
func HandlerPostArticles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := GetCtxUsecase(ctx)
		body := struct {
			Title       string   `json:"title"`
			Description string   `json:"description"`
			PublishedAt int64    `json:"publishedAt"`
			Tags        []string `json:"tags"`
		}{}
		if err := ctx.BindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				err.Error(),
			)
			return
		}
		article := model.Article{}
		if err := u.CreateArticle(ctx, &article); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				err.Error(),
			)
			return
		}
		now := time.Now()
		article.CreatedAt = now.Unix()
		article.UpdatedAt = now.Unix()
		ctx.JSON(http.StatusCreated, NewResponseArticle(&article))
		return
	}
}

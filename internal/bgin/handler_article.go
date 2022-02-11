package bgin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/internal/entity"
	"github.com/suzuito/blog1-go/internal/usecase"
	"github.com/suzuito/common-go/cgin"
)

// HandlerGetArticles ...
func HandlerGetArticles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()
		u := getCtxUsecase(ctx)
		n := cgin.DefaultQueryAsInt(ctx, "n", 10)
		cursorPublishedAt := cgin.DefaultQueryAsInt64(ctx, "cursor_published_at", now.Unix())
		cursorTitle := ctx.DefaultQuery("cursor_title", "")
		order := usecase.CursorOrder(ctx.DefaultQuery("order", string(usecase.CursorOrderDesc)))
		articles := []entity.Article{}
		if err := u.GetArticles(ctx, cursorPublishedAt, cursorTitle, order, n, &articles); err != nil {
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
		article := getCtxArticle(ctx)
		ctx.JSON(
			http.StatusOK,
			NewResponseArticle(article),
		)
	}
}

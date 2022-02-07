package bgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/internal/entity/model"
	"github.com/suzuito/blog1-go/internal/usecase"
	"golang.org/x/xerrors"
)

// MiddlewareGetArticle ...
func MiddlewareGetArticle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := getCtxUsecase(ctx)
		articleID := model.ArticleID(ctx.Param("articleID"))
		article := model.Article{}
		if err := u.GetArticle(ctx, articleID, &article); err != nil {
			if xerrors.Is(err, usecase.ErrNotFound) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, NewResponseError(err))
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewResponseError(err))
			return
		}
		setCtxArticle(ctx, &article)
	}
}

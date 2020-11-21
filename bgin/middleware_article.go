package bgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/entity/model"
	"github.com/suzuito/blog1-go/usecase"
	"golang.org/x/xerrors"
)

// MiddlewareGetArticle ...
func MiddlewareGetArticle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := GetCtxUsecase(ctx)
		articleID := model.ArticleID(ctx.Param("articleID"))
		article := model.Article{}
		err := u.GetArticle(articleID, &article)
		if xerrors.Is(err, usecase.ErrNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, NewResponseError(err))
			return
		}
		SetCtxArticle(ctx, &article)
	}
}

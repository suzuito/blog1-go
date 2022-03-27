package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/blog1-go/pkg/usecase"
	"golang.org/x/xerrors"
)

// HTMLMiddlewareGetArticle ...
func HTMLMiddlewareGetArticle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := getCtxUsecase(ctx)
		articleID := entity.ArticleID(ctx.Param("articleID"))
		article := entity.Article{}
		if err := u.GetArticle(ctx, articleID, &article); err != nil {
			if xerrors.Is(err, usecase.ErrNotFound) {
				ctx.Abort()
				html404(ctx)
				return
			}
			ctx.Abort()
			html500(ctx, err)
			return
		}
		setCtxArticle(ctx, &article)
	}
}

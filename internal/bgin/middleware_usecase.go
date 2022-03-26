package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

// MiddlewareUsecase ...
func MiddlewareUsecase(gdeps *inject.GlobalDepends) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := usecase.NewImpl(
			gdeps.DB,
			gdeps.Storage,
			gdeps.MDConverter,
			gdeps.HTMLMediaFetcher,
			gdeps.HTMLEditor,
			gdeps.HTMLTOCExtractor,
		)
		setCtxUsecase(ctx, u)
		ctx.Next()
	}
}

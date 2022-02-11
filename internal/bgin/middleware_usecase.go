package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/internal/inject"
	"github.com/suzuito/blog1-go/internal/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

// MiddlewareUsecase ...
func MiddlewareUsecase(env *setting.Environment, gdeps *inject.GlobalDepends) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := usecase.NewImpl(
			gdeps.DB,
			gdeps.Storage,
			gdeps.MDConverter,
		)
		setCtxUsecase(ctx, u)
		ctx.Next()
	}
}

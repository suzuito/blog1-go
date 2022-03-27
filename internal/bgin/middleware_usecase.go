package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

// MiddlewareUsecase ...
func MiddlewareUsecase(u usecase.Usecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		setCtxUsecase(ctx, u)
		ctx.Next()
	}
}

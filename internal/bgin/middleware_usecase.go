package bgin

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/internal/inject"
	"github.com/suzuito/blog1-go/internal/setting"
	"github.com/suzuito/blog1-go/internal/usecase"
)

// MiddlewareUsecase ...
func MiddlewareUsecase(env *setting.Environment, gdeps *inject.GlobalDepends) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cdeps, ccloseFunc, err := inject.NewContextDepends(ctx, env)
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
		defer ccloseFunc()
		u := usecase.NewImpl(
			env,
			cdeps.DB,
			cdeps.Storage,
			gdeps.MDConverter,
		)
		setCtxUsecase(ctx, u)
		ctx.Next()
	}
}

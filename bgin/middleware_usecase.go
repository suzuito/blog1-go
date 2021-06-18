package bgin

import (
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/application"
	"github.com/suzuito/blog1-go/bgcp/fdb"
	"github.com/suzuito/blog1-go/setting"
	"github.com/suzuito/blog1-go/usecase"
)

// MiddlewareUsecase ...
func MiddlewareUsecase(app *application.Application, env *setting.Environment) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fcli, err := firestore.NewClient(ctx, app.GCPProjectID)
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
		defer fcli.Close()
		u := usecase.NewImpl(
			env,
			nil,
			fdb.NewClient(fcli),
			nil,
			nil,
		)
		setCtxUsecase(ctx, u)
		ctx.Next()
	}
}

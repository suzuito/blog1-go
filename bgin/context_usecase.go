package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/usecase"
)

func getCtxUsecase(ctx *gin.Context) usecase.Usecase {
	u, exists := ctx.Get("usecase")
	if !exists {
		return nil
	}
	uu, _ := u.(usecase.Usecase)
	return uu
}

func setCtxUsecase(ctx *gin.Context, u usecase.Usecase) {
	ctx.Set("usecase", u)
}

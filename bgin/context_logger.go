package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/xlogging"
)

func getCtxLogger(ctx *gin.Context) xlogging.Logger {
	u, exists := ctx.Get("logger")
	if !exists {
		return nil
	}
	uu, _ := u.(xlogging.Logger)
	return uu
}

func setCtxLogger(ctx *gin.Context, v xlogging.Logger) {
	ctx.Set("logger", v)
}

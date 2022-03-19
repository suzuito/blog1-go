package bgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/setting"
)

func html404(ctx *gin.Context, env *setting.Environment) {
	ctx.HTML(
		http.StatusNotFound,
		"pc_404.html",
		gin.H{
			"Global": htmlGlobal(env),
		},
	)
}

func html500(ctx *gin.Context, env *setting.Environment) {
	ctx.HTML(
		http.StatusInternalServerError,
		"pc_500.html",
		gin.H{
			"Global": htmlGlobal(env),
		},
	)
}

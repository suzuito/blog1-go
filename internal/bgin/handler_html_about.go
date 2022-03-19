package bgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/setting"
)

// HandlerHTMLGetAbout ...
func HandlerHTMLGetAbout(
	env *setting.Environment,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"pc_about.html",
			gin.H{
				"Global": htmlGlobal(env),
			},
		)
	}
}

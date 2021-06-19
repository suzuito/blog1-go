package bgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerGetSitemapXML ...
func HandlerGetSitemapXML() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := getCtxUsecase(ctx)
		origin := ctx.DefaultQuery("origin", "http://localhost:4200")
		b, err := u.GenerateBlogSiteMap(ctx, origin)
		if err != nil {
			ctx.AbortWithStatus(
				http.StatusInternalServerError,
			)
			return
		}
		ctx.Header("Content-type", "application/xml")
		body, err := b.Marshal()
		if err != nil {
			ctx.AbortWithStatus(
				http.StatusInternalServerError,
			)
			return
		}
		ctx.String(http.StatusOK, body)
	}
}

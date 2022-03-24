package bgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/setting"
)

// HandlerGetSitemapXML ...
func HandlerGetSitemapXML() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := getCtxUsecase(ctx)
		b, err := u.GenerateBlogSiteMap(ctx, setting.E.SiteOrigin)
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

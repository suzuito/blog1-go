package bgin

import (
	"encoding/xml"
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
		bb, err := xml.MarshalIndent(b, "", "    ")
		if err != nil {
			ctx.AbortWithStatus(
				http.StatusInternalServerError,
			)
			return
		}
		c := string(bb)
		c = `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + c
		ctx.String(http.StatusOK, c)
	}
}

package bgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerHTMLGetAbout ...
func HandlerHTMLGetAbout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"pc_about.html",
			newTmplVar(
				newTmplVarMeta(
					"ブログサイト管理人の自己紹介",
				),
				newTmplVarLink(
					getPageURL(ctx),
				),
				newTmplVarOGP(
					"ブログサイト管理人の自己紹介",
					"ブログサイト管理人の自己紹介",
					"website",
					getPageURL(ctx),
					"",
				),
				[]tmplVarLDJSON{
					newTmplVarLDJSONWebSite(
						getPageURL(ctx),
						"ブログサイト管理人の自己紹介",
						"ブログサイト管理人の自己紹介",
					),
				},
				map[string]interface{}{},
			),
		)
	}
}

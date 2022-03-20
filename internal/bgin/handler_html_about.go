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
			newTmplVar(
				env,
				newTmplVarMeta(
					"ブログサイト管理人の自己紹介",
				),
				newTmplVarLink(
					getPageURL(ctx, env),
				),
				newTmplVarOGP(
					"ブログサイト管理人の自己紹介",
					"ブログサイト管理人の自己紹介",
					"website",
					getPageURL(ctx, env),
					"",
				),
				[]tmplVarLDJSON{
					newTmplVarLDJSONWebSite(
						getPageURL(ctx, env),
						"ブログサイト管理人の自己紹介",
						"ブログサイト管理人の自己紹介",
					),
				},
				map[string]interface{}{},
			),
		)
	}
}

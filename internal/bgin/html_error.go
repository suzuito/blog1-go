package bgin

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func html404(ctx *gin.Context) {
	ctx.HTML(
		http.StatusNotFound,
		"pc_404.html",
		newTmplVar(
			newTmplVarMeta(
				metaSiteDescription,
			),
			newTmplVarLink(
				getPageURL(ctx),
			),
			newTmplVarOGP(
				"404エラーページ",
				"404エラーページ",
				"website",
				getPageURL(ctx),
				"",
			),
			[]tmplVarLDJSON{
				newTmplVarLDJSONWebSite(
					getPageURL(ctx),
					"404エラーページ",
					"404エラーページ",
				),
			},
			map[string]interface{}{},
		),
	)
}

func html500(ctx *gin.Context, err error) {
	ctx.HTML(
		http.StatusInternalServerError,
		"pc_500.html",
		newTmplVar(
			newTmplVarMeta(
				metaSiteDescription,
			),
			newTmplVarLink(
				getPageURL(ctx),
			),
			newTmplVarOGP(
				"500エラーページ",
				"500エラーページ",
				"website",
				getPageURL(ctx),
				"",
			),
			[]tmplVarLDJSON{
				newTmplVarLDJSONWebSite(
					getPageURL(ctx),
					"500エラーページ",
					"500エラーページ",
				),
			},
			map[string]interface{}{},
		),
	)

	if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
		hub.WithScope(func(scope *sentry.Scope) {
			hub.CaptureException(err)
		})
	}
}

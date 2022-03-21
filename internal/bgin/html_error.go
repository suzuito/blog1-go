package bgin

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/setting"
)

func html404(ctx *gin.Context, env *setting.Environment) {
	ctx.HTML(
		http.StatusNotFound,
		"pc_404.html",
		newTmplVar(
			env,
			newTmplVarMeta(
				metaSiteDescription,
			),
			newTmplVarLink(
				getPageURL(ctx, env),
			),
			newTmplVarOGP(
				"404エラーページ",
				"404エラーページ",
				"website",
				getPageURL(ctx, env),
				"",
			),
			[]tmplVarLDJSON{
				newTmplVarLDJSONWebSite(
					getPageURL(ctx, env),
					"404エラーページ",
					"404エラーページ",
				),
			},
			map[string]interface{}{},
		),
	)
}

func html500(ctx *gin.Context, env *setting.Environment, err error) {
	ctx.HTML(
		http.StatusInternalServerError,
		"pc_500.html",
		newTmplVar(
			env,
			newTmplVarMeta(
				metaSiteDescription,
			),
			newTmplVarLink(
				getPageURL(ctx, env),
			),
			newTmplVarOGP(
				"500エラーページ",
				"500エラーページ",
				"website",
				getPageURL(ctx, env),
				"",
			),
			[]tmplVarLDJSON{
				newTmplVarLDJSONWebSite(
					getPageURL(ctx, env),
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

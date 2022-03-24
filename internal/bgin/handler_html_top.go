package bgin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

// HandlerHTMLGetTop ...
func HandlerHTMLGetTop() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now().Unix()
		u := getCtxUsecase(ctx)
		articles := []entity.Article{}
		if err := u.GetArticles(ctx, now, "", usecase.CursorOrderDesc, 3, &articles); err != nil {
			html500(ctx, err)
			return
		}
		ctx.HTML(
			http.StatusOK,
			"pc_top.html",
			newTmplVar(
				newTmplVarMeta(
					metaSiteDescription,
				),
				newTmplVarLink(
					getPageURL(ctx),
				),
				newTmplVarOGP(
					metaSiteName,
					metaSiteDescription,
					"website",
					getPageURL(ctx),
					"",
				),
				[]tmplVarLDJSON{
					newTmplVarLDJSONWebSite(
						getPageURL(ctx),
						metaSiteName,
						metaSiteDescription,
					),
				},
				map[string]interface{}{
					"Articles": articles,
				},
			),
		)
	}
}

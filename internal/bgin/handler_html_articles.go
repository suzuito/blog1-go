package bgin

import (
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/blog1-go/pkg/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
	"github.com/suzuito/common-go/cgin"
)

// HandlerHTMLGetArticles ...
func HandlerHTMLGetArticles(
	env *setting.Environment,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()
		u := getCtxUsecase(ctx)
		n := cgin.DefaultQueryAsInt(ctx, "n", 10)
		cursorPublishedAt := cgin.DefaultQueryAsInt64(ctx, "cursor_published_at", now.Unix())
		cursorTitle := ctx.DefaultQuery("cursor_title", "")
		order := usecase.CursorOrder(ctx.DefaultQuery("order", string(usecase.CursorOrderDesc)))
		articles := []entity.Article{}
		if err := u.GetArticles(ctx, cursorPublishedAt, cursorTitle, order, n, &articles); err != nil {
			html500(ctx, env, err)
			return
		}
		sort.Slice(articles, func(i, j int) bool {
			return articles[i].PublishedAt > articles[j].PublishedAt
		})
		nextPublishedAt := int64(0)
		nextTitle := ""
		if len(articles) >= n {
			nextPublishedAt = articles[len(articles)-1].PublishedAt
			nextTitle = articles[len(articles)-1].Title
		}
		prevPublishedAt := int64(0)
		prevTitle := ""
		if len(articles) > 0 {
			prevPublishedAt = articles[0].PublishedAt
			prevTitle = articles[0].Title
		}
		ctx.HTML(
			http.StatusOK,
			"pc_articles.html",
			newTmplVar(
				env,
				newTmplVarMeta(
					"記事一覧",
				),
				newTmplVarLink(
					getPageURL(ctx, env),
				),
				newTmplVarOGP(
					"記事一覧",
					"記事一覧",
					"article",
					getPageURL(ctx, env),
					"",
				),
				[]tmplVarLDJSON{
					newTmplVarLDJSONWebSite(
						getPageURL(ctx, env),
						"記事一覧",
						"記事一覧",
					),
				},
				map[string]interface{}{
					"Articles":        articles,
					"NextPublishedAt": nextPublishedAt,
					"NextTitle":       nextTitle,
					"PrevPublishedAt": prevPublishedAt,
					"PrevTitle":       prevTitle,
				},
			),
		)
	}
}

package bgin

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

// HandlerHTMLGetArticle ...
func HandlerHTMLGetArticle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		article := getCtxArticle(ctx)
		u := getCtxUsecase(ctx)
		content := []byte{}
		if err := u.GetArticleHTML(ctx, article.ID, &content); err != nil {
			if errors.Is(err, usecase.ErrNotFound) {
				html404(ctx)
				return
			}
			html500(ctx, err)
			return
		}
		imageURLs := []string{
			getAvatarURL(),
		}
		for _, img := range imageURLs {
			imageURLs = append(imageURLs, img)
		}
		imageURL := ""
		if len(imageURLs) > 0 {
			imageURL = imageURLs[0]
		}
		buf := bytes.NewBufferString("")
		if err := tmplArticle.ExecuteTemplate(
			buf,
			"pc_article.html",
			newTmplVar(
				newTmplVarMeta(
					article.Description,
				),
				newTmplVarLink(
					getPageURL(ctx),
				),
				newTmplVarOGP(
					article.Title,
					article.Description,
					"article",
					getPageURL(ctx),
					imageURL,
				),
				[]tmplVarLDJSON{
					newTmplVarLDJSONArticle(
						article.Description,
						article.Description,
						article.CreatedAtAsTime(),
						imageURLs,
						"otiuzu",
						getAboutPageURL(),
					),
				},
				map[string]interface{}{
					"Article": article,
				},
			),
		); err != nil {
			html500(ctx, err)
			return
		}
		body := strings.Replace(buf.String(), "__QSW#$%FG_CONTENT__", string(content), -1)
		ctx.Header("Content-type", "text/html; charset=UTF-8")
		ctx.String(http.StatusOK, body)
	}
}

package bgin

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/setting"
	"golang.org/x/xerrors"
)

var tmplArticle *template.Template

func init() {
	var err error
	tmplArticle, err = template.New("hoge").ParseGlob("data/template/*.html")
	if err != nil {
		panic(xerrors.Errorf("cannot new template : %+v", err))
	}
}

// HandlerHTMLGetArticle ...
func HandlerHTMLGetArticle(
	env *setting.Environment,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		article := getCtxArticle(ctx)
		u := getCtxUsecase(ctx)
		content := []byte{}
		if err := u.GetArticleHTML(ctx, article.ID, &content); err != nil {
			html404(ctx, env)
			return
		}
		buf := bytes.NewBufferString("")
		if err := tmplArticle.ExecuteTemplate(
			buf,
			"pc_article.html",
			gin.H{
				"Global":  htmlGlobal(env),
				"Article": article,
				"LDJSON": map[string]interface{}{
					"@context":      "https://schema.org",
					"@type":         "Article",
					"headline":      article.Description,
					"datePublished": article.PublishedAtAsTime().Format(time.RFC3339),
					"image": (func() []string {
						a := []string{}
						for _, img := range article.Images {
							a = append(a, img.URL)
						}
						return a
					})(),
				},
			},
		); err != nil {
			html500(ctx, env)
			return
		}
		body := strings.Replace(buf.String(), "__QSW#$%FG_CONTENT__", string(content), -1)
		ctx.Header("Content-type", "text/html; charset=UTF-8")
		ctx.String(http.StatusOK, body)
	}
}

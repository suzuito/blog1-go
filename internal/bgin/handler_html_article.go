package bgin

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/setting"
	"golang.org/x/xerrors"
)

// HandlerHTMLGetArticle ...
func HandlerHTMLGetArticle(
	env *setting.Environment,
) gin.HandlerFunc {
	var once sync.Once
	var tmplArticle *template.Template
	return func(ctx *gin.Context) {
		var errTmpl error
		once.Do(func() {
			var err error
			tmplArticle, errTmpl = template.New("hoge").ParseGlob(fmt.Sprintf("%s/*.html", env.DirPathTemplate))
			if err != nil {
				errTmpl = xerrors.Errorf("cannot new template : %+v", err)
			}
		})
		if errTmpl != nil {
			html500(ctx, env, errTmpl)
			return
		}
		article := getCtxArticle(ctx)
		u := getCtxUsecase(ctx)
		content := []byte{}
		if err := u.GetArticleHTML(ctx, article.ID, &content); err != nil {
			html404(ctx, env)
			return
		}
		imageURLs := []string{
			getAvatarURL(env),
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
				env,
				newTmplVarMeta(
					article.Description,
				),
				newTmplVarLink(
					getPageURL(ctx, env),
				),
				newTmplVarOGP(
					article.Title,
					article.Description,
					"article",
					getPageURL(ctx, env),
					imageURL,
				),
				[]tmplVarLDJSON{
					newTmplVarLDJSONArticle(
						article.Description,
						article.Description,
						article.CreatedAtAsTime(),
						imageURLs,
						"otiuzu",
						getAboutPageURL(env),
					),
				},
				map[string]interface{}{
					"Article": article,
				},
			),
		); err != nil {
			html500(ctx, env, err)
			return
		}
		body := strings.Replace(buf.String(), "__QSW#$%FG_CONTENT__", string(content), -1)
		ctx.Header("Content-type", "text/html; charset=UTF-8")
		ctx.String(http.StatusOK, body)
	}
}

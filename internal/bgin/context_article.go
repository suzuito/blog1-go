package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/internal/entity/model"
)

func getCtxArticle(ctx *gin.Context) *model.Article {
	u, exists := ctx.Get("article")
	if !exists {
		return nil
	}
	uu, _ := u.(*model.Article)
	return uu
}

func setCtxArticle(ctx *gin.Context, a *model.Article) {
	ctx.Set("article", a)
}

package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/entity"
)

func getCtxArticle(ctx *gin.Context) *entity.Article {
	u, exists := ctx.Get("article")
	if !exists {
		return nil
	}
	uu, _ := u.(*entity.Article)
	return uu
}

func setCtxArticle(ctx *gin.Context, a *entity.Article) {
	ctx.Set("article", a)
}

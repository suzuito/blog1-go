package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/internal/entity/model"
)

func getCtxAdminAuth(ctx *gin.Context) *model.AdminAuth {
	u, exists := ctx.Get("admin_auth")
	if !exists {
		return nil
	}
	uu, _ := u.(*model.AdminAuth)
	return uu
}

func setCtxAdminAuth(ctx *gin.Context, a *model.AdminAuth) {
	ctx.Set("admin_auth", a)
}

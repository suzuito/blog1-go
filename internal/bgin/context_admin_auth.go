package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/internal/entity"
)

func getCtxAdminAuth(ctx *gin.Context) *entity.AdminAuth {
	u, exists := ctx.Get("admin_auth")
	if !exists {
		return nil
	}
	uu, _ := u.(*entity.AdminAuth)
	return uu
}

func setCtxAdminAuth(ctx *gin.Context, a *entity.AdminAuth) {
	ctx.Set("admin_auth", a)
}

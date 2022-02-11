package bgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/internal/entity"
)

var headerNameAdminAuth = "X-Admin-Auth"

// MiddlewareAdminAuth ...
func MiddlewareAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := getCtxUsecase(ctx)
		adminAuthHeader := ctx.GetHeader(headerNameAdminAuth)
		adminAuth := entity.AdminAuth{}
		err := u.GetAdminAuth(ctx, adminAuthHeader, &adminAuth)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ResponseError{
				Detail: "Unauthorization",
			})
			return
		}
		setCtxAdminAuth(ctx, &adminAuth)
	}
}

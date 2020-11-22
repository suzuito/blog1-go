package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/application"
)

// MiddlewareGetTag ...
func MiddlewareGetTag(app *application.Application) gin.HandlerFunc {
	return func(context *gin.Context) {}
}

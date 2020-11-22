package bgin

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/application"
)

// MiddlewareLogger ...
func MiddlewareLogger(app *application.Application) gin.HandlerFunc {
	return func(context *gin.Context) {}
}

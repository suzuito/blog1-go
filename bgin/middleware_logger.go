package bgin

import (
	"github.com/gin-gonic/gin"
)

// MiddlewareLogger ...
func MiddlewareLogger() gin.HandlerFunc {
	return func(context *gin.Context) {}
}

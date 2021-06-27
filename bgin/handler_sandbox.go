package bgin

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandlerGetSandbox() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		statusString := ctx.Query("status")
		status := http.StatusOK
		if statusString != "" {
			var err error
			status, err = strconv.Atoi(statusString)
			if err != nil {
				status = http.StatusOK
			}
		}
		fmt.Fprint(os.Stderr, "hoge\n")
		ctx.Status(
			status,
		)
	}
}

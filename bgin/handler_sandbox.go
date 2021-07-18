package bgin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/xlogging"
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
		logger := getCtxLogger(ctx)
		logger.Payloadf(xlogging.SeverityDebug, "hoge1")
		logger.Payloadf(xlogging.SeverityInfo, "hoge2")
		logger.Payloadf(xlogging.SeverityError, "hoge3")
		ctx.Status(
			status,
		)
	}
}

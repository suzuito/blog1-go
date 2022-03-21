package bgin

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/suzuito/blog1-go/pkg/setting"
)

func HandlerHTMLGetSandbox(
	env *setting.Environment,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Error().Err(fmt.Errorf("err1")).Msgf("dummy error001")
		log.Info().Interface("data", map[string]string{"hoge": "fuga"}).Msgf("dummy info")
		log.Warn().Interface("data", map[string]string{"hoge": "fuga"}).Msgf("dummy warn")
		// Report to Cloud Logging
		payload := map[string]interface{}{
			"level": "error", // FIXME required?
			"serviceContext": map[string]interface{}{
				"service": "blog", // required
				"version": "1.0.1",
			},
			"message": "Dummy error is occured 001", // required
			"context": map[string]interface{}{
				"httpRequest": map[string]interface{}{
					"method":    c.Request.Method,
					"url":       c.Request.URL.String(),
					"userAgent": c.Request.UserAgent(),
					"referrer":  c.Request.Referer(),
				},
				// "reportLocation": map[string]interface{}{ // FIXME required?
				// 	"filePath":     "hoge.go",
				// 	"fileNumber":   102,
				// 	"functionName": "fuga",
				// },
			},
		}
		payloadBytes, _ := json.MarshalIndent(payload, "", " ")
		fmt.Println(string(payloadBytes))
	}
}

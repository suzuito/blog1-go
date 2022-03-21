package bgin

import (
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
		serviceContext := map[string]string{
			"service": "test001",
		}
		// Report to Cloud Logging
		payload := map[string]interface{}{
			"serviceContext": map[string]interface{}{
				"service": "blog",
				"version": "1.0.0",
			},
			"message": "Dummy error is occured",
			"context": map[string]interface{}{
				"reportLocation": map[string]interface{}{
					"filePath":     "hoge.go",
					"fileNumber":   101,
					"functionName": "fuga",
				},
			},
		}
		fmt.Println(payload)
	}
}

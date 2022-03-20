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
		log.Error().Err(fmt.Errorf("err1")).Msgf("dummy error")
		log.Info().Interface("data", map[string]string{"hoge": "fuga"}).Msgf("dummy info")
		log.Warn().Interface("data", map[string]string{"hoge": "fuga"}).Msgf("dummy warn")
	}
}

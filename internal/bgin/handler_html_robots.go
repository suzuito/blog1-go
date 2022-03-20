package bgin

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/setting"
)

func HandlerHTMLGetRobots(env *setting.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := url.Parse(env.SiteOrigin)
		if err != nil {
			u = &url.URL{
				Scheme: "http",
				Host:   "localhost:8888",
			}
		}
		u.Path = "sitemap.xml"
		if env.Env == "godzilla" {
			c.String(http.StatusOK, strings.Join(
				[]string{
					"Sitemap: " + u.String(),
				}, "\n"),
			)
			return
		}
		c.String(http.StatusOK, strings.Join(
			[]string{
				"Sitemap: " + u.String(),
				"User-agent: *",
				"Disallow: /",
			}, "\n"),
		)
	}
}

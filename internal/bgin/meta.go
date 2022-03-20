package bgin

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/pkg/setting"
)

var (
	metaSiteDescription = "ブログサイト。"
	metaSiteName        = "otiuzu pages"
)

func getPageURL(ctx *gin.Context, env *setting.Environment) string {
	u, _ := url.Parse(env.SiteOrigin)
	if ctx.Request != nil {
		u.Path = ctx.Request.URL.Path
	}
	return u.String()
}

func getAboutPageURL(env *setting.Environment) string {
	u, _ := url.Parse(env.SiteOrigin)
	u.Path = "about"
	return u.String()
}

func getAvatarURL(env *setting.Environment) string {
	u, _ := url.Parse(env.SiteOrigin)
	u.Path = "asset/avatar.jpg"
	return u.String()
}

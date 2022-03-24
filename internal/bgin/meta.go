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

func getPageURL(ctx *gin.Context) string {
	u, _ := url.Parse(setting.E.SiteOrigin)
	if ctx.Request != nil {
		u.Path = ctx.Request.URL.Path
	}
	return u.String()
}

func getAboutPageURL() string {
	u, _ := url.Parse(setting.E.SiteOrigin)
	u.Path = "about"
	return u.String()
}

func getAvatarURL() string {
	u, _ := url.Parse(setting.E.SiteOrigin)
	u.Path = "asset/avatar.jpg"
	return u.String()
}

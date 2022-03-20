package bgin

import (
	"time"

	"github.com/suzuito/blog1-go/pkg/setting"
)

type tmplVarMeta map[string]interface{}

func newTmplVarMeta(
	description string,
) tmplVarMeta {
	return tmplVarMeta{
		"Description": description,
	}
}

type tmplVarLink map[string]interface{}

func newTmplVarLink(
	canonical string,
) tmplVarLink {
	return tmplVarLink{
		"Canonical": canonical,
	}
}

type tmplVarOGP map[string]interface{}

func newTmplVarOGP(
	title string,
	description string,
	typeOgp string,
	url string,
	image string,
) tmplVarOGP {
	r := tmplVarOGP{
		"Title":       title,
		"Description": description,
		"Local":       "ja_JP",
		"Type":        typeOgp,
		"URL":         url,
		"SiteName":    metaSiteName,
	}
	if image != "" {
		r["Image"] = image
	}
	return r
}

type tmplVarLDJSON map[string]interface{}

func newTmplVarLDJSONWebSite(
	mainEntityOfPage string,
	headline string,
	description string,
) tmplVarLDJSON {
	r := tmplVarLDJSON{
		"@context":         "https://schema.org",
		"@type":            "WebSite",
		"mainEntityOfPage": mainEntityOfPage,
		"headline":         headline,
		"description":      description,
	}

	return r
}

func newTmplVarLDJSONArticle(
	headline string,
	description string,
	datePublished time.Time,
	image string,
) tmplVarLDJSON {
	r := tmplVarLDJSON{
		"@context":      "https://schema.org",
		"@type":         "Article",
		"headline":      headline,
		"description":   description,
		"datePublished": datePublished.Format(time.RFC3339),
	}
	if image != "" {
		r["image"] = image
	}
	return r
}

func newTmplVar(
	env *setting.Environment,
	meta tmplVarMeta,
	link tmplVarLink,
	ogp tmplVarOGP,
	ldjson []tmplVarLDJSON,
	other map[string]interface{},
) map[string]interface{} {
	ret := map[string]interface{}{
		"_Global": map[string]interface{}{
			"Title": "otiuzu pages",
			"GA":    env.GA,
		},
		"_LDJSON": ldjson,
		"_Link":   link,
		"_Meta":   meta,
		"_OGP":    ogp,
	}
	for k, v := range other {
		ret[k] = v
	}
	return ret
}

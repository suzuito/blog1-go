package usecase

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/url"
	"time"

	"github.com/suzuito/blog1-go/entity/model"
	"golang.org/x/xerrors"
)

type xmlURLSet struct {
	XMLName           xml.Name `xml:"urlset"`
	URLs              []xmlURL `xml:"url"`
	XMLNSXsi          string   `xml:"xmlns:xsi,attr"`
	XMLNS             string   `xml:"xmlns,attr"`
	XsiSchemaLocation string   `xml:"xsi:schemaLocation,attr"`
}

type xmlURL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
	Lastmod string   `xml:"lastmod"`
}

func newXMLURLFromArticle(a *model.Article, origin string) *xmlURL {
	mod := time.Unix(a.PublishedAt, 0).Format("2006-01-02")
	return &xmlURL{
		Loc:     fmt.Sprintf("%s/blog/%s", origin, url.QueryEscape(string(a.ID))),
		Lastmod: mod,
	}
}

// GenerateBlogSiteMap ...
func (u *Impl) GenerateBlogSiteMap(ctx context.Context, origin string) (string, error) {
	urls := xmlURLSet{
		URLs: []xmlURL{},
	}
	// Articles
	cursorPublishedAt := int64(0)
	cursorTitle := ""
	for {
		articles := []model.Article{}
		if err := u.db.GetArticles(ctx, cursorPublishedAt, cursorTitle, CursorOrderAsc, 100, &articles); err != nil {
			return "", xerrors.Errorf("Cannot get articles : %w", err)
		}
		if len(articles) <= 0 {
			break
		}
		for _, article := range articles {
			x := newXMLURLFromArticle(&article, origin)
			urls.URLs = append(urls.URLs, *x)
		}
		last := articles[len(articles)-1]
		cursorPublishedAt = last.PublishedAt
		cursorTitle = last.Title
	}

	urls.URLs = append(urls.URLs, xmlURL{
		Lastmod: "2020-11-01",
		Loc:     fmt.Sprintf("%s", origin),
	})
	urls.URLs = append(urls.URLs, xmlURL{
		Lastmod: "2020-11-01",
		Loc:     fmt.Sprintf("%s/blog", origin),
	})

	urls.XMLNSXsi = "http://www.w3.org/2001/XMLSchema-instance"
	urls.XMLNS = "http://www.sitemaps.org/schemas/sitemap/0.9"
	urls.XsiSchemaLocation = "http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd"

	b, err := xml.MarshalIndent(urls, "", "    ")
	if err != nil {
		return "", xerrors.Errorf("Cannot marshal xml : %w", err)
	}

	c := string(b)
	c = `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + c
	return c, nil
}

package local

import (
	"bytes"
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday/v2"
	"github.com/suzuito/blog1-go/entity/model"
	"golang.org/x/xerrors"
)

// BlackFridayMDConverter ...
type BlackFridayMDConverter struct{}

// Convert ...
func (c *BlackFridayMDConverter) Convert(
	ctx context.Context,
	article *model.Article,
	raw []byte,
	converted *[]byte,
) error {
	*converted = blackfriday.Run(
		raw,
		blackfriday.WithRenderer(
			blackfriday.NewHTMLRenderer(
				blackfriday.HTMLRendererParameters{
					Flags: blackfriday.TOC | blackfriday.HrefTargetBlank,
				},
			),
		),
	)
	converted1, err := convertAfterConvert(*converted)
	if err != nil {
		return xerrors.Errorf("Cannot convertAfterConvert : %w", err)
	}
	*converted = []byte(converted1)
	return nil
}

func convertAfterConvert(body []byte) (string, error) {
	d, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return "", xerrors.Errorf("Cannot new goquery : %w", err)
	}
	d.Find("pre").Each(func(i int, s *goquery.Selection) {
		s.SetAttr("class", "code-block")
		s.SetAttr("style", "width: 100%; overflow: scroll;")
	})
	returned, err := d.Html()
	if err != nil {
		return "", xerrors.Errorf("Cannot create goquery html : %w", err)
	}
	returned = strings.Replace(returned, "<html><head></head><body>", "", 1)
	returned = strings.Replace(returned, "</body></html>", "", 1)
	return returned, nil
}

package bhtml

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

type Editor struct {
}

func (b *Editor) ModifyHTML(ctx context.Context, src string, dst *string) error {
	tmp := strings.Replace(src, "<html><head></head><body>", "", 1)
	tmp = strings.Replace(tmp, "</body></html>", "", 1)
	d, err := goquery.NewDocumentFromReader(strings.NewReader(tmp))
	if err != nil {
		return errors.Wrapf(err, "cannot convert to html")
	}
	d.Find("pre").Each(func(i int, s *goquery.Selection) {
		s.SetAttr("class", "code-block")
		s.SetAttr("style", "width: 100%; overflow: scroll;")
	})
	d.Find("img.md-image").Each(func(i int, s *goquery.Selection) {
		s.SetAttr("style", "width: 100%;")
	})
	*dst, err = d.Html()
	if err != nil {
		return errors.Wrapf(err, "cannot HTML")
	}
	return nil
}

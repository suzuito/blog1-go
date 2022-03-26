package bhtml

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/entity"
)

type TOCExtractor struct{}

func (e *TOCExtractor) Extract(
	src string,
	idx *[]entity.ArticleIndex,
) error {
	d, err := goquery.NewDocumentFromReader(strings.NewReader(src))
	if err != nil {
		return errors.Wrapf(err, "cannot convert to html")
	}
	d.Find(".md-heading").Each(func(i int, s *goquery.Selection) {
		toc := entity.ArticleIndex{
			Name:  s.Text(),
			ID:    s.AttrOr("id", ""),
			Level: entity.NewArticleIndexLevel(goquery.NodeName(s)),
		}
		*idx = append(*idx, toc)
	})
	return nil
}

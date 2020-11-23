package local

import (
	"context"

	"github.com/russross/blackfriday/v2"
	"github.com/suzuito/blog1-go/entity/model"
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
					Flags: blackfriday.TOC,
				},
			),
		),
	)
	return nil
}

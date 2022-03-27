package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/pkg/entity"
)

type HTMLEditor interface {
	ModifyHTML(ctx context.Context, src string, dst *string) error
}

type HTMLMediaFetcher interface {
	Fetch(
		ctx context.Context,
		src string,
		images *[]entity.ArticleImage,
	) error
}

type HTMLTOCExtractor interface {
	Extract(
		src string,
		idx *[]entity.ArticleIndex,
	) error
}

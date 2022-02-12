package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/pkg/entity"
)

// ArticleReader ...
// :Deprecated
type ArticleReader interface {
	Walk(ctx context.Context, each func(article *entity.Article, raw []byte) error) error
	Close() error
}

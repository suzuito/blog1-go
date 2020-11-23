package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/entity/model"
)

// ArticleReader ...
type ArticleReader interface {
	Walk(ctx context.Context, each func(article *model.Article, raw []byte) error) error
	Close() error
}

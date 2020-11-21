package usecase

import (
	"context"
	"fmt"

	"github.com/suzuito/blog1-go/entity/model"
)

var (
	// ErrNotFound ...
	ErrNotFound = fmt.Errorf("Not found")
)

// DB ...
type DB interface {
	GetArticles(
		ctx context.Context,
		startPublishedAt int64,
		n int,
		articles *[]model.Article,
	) error
	GetArticle(
		ctx context.Context,
		articleID model.ArticleID,
		article *model.Article,
	) error
}

package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/internal/entity/model"
)

// CursorOrder ...
type CursorOrder string

var (
	// CursorOrderAsc ...
	CursorOrderAsc CursorOrder = "asc"
	// CursorOrderDesc ...
	CursorOrderDesc CursorOrder = "desc"
)

// DB ...
type DB interface {
	GetArticles(
		ctx context.Context,
		cursorPublishedAt int64,
		cursorTitle string,
		order CursorOrder,
		n int,
		articles *[]model.Article,
	) error
	GetArticle(
		ctx context.Context,
		articleID model.ArticleID,
		article *model.Article,
	) error
	SetArticle(
		ctx context.Context,
		article *model.Article,
	) error
}

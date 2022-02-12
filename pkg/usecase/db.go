package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/pkg/entity"
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
		articles *[]entity.Article,
	) error
	GetArticle(
		ctx context.Context,
		articleID entity.ArticleID,
		article *entity.Article,
	) error
	SetArticle(
		ctx context.Context,
		article *entity.Article,
	) error
	DeleteArticle(
		ctx context.Context,
		articleID entity.ArticleID,
	) error
}

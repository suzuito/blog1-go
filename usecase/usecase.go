package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/entity/model"
)

// Usecase ...
type Usecase interface {
	SyncArticles(
		ctx context.Context,
		source ArticleReader,
	) error

	GetArticle(
		ctx context.Context,
		id model.ArticleID,
		article *model.Article,
	) error
	GetArticles(
		ctx context.Context,
		cursorPublishedAt int64,
		cursorTitle string,
		order CursorOrder,
		n int,
		articles *[]model.Article,
	) error
}

// Impl ...
type Impl struct {
	db DB
}

// NewImpl ...
func NewImpl(db DB) *Impl {
	return &Impl{
		db: db,
	}
}

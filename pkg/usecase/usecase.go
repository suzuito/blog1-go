package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/internal/cmarkdown"
	"github.com/suzuito/blog1-go/pkg/entity"
)

// Usecase ...
type Usecase interface {
	GetArticle(
		ctx context.Context,
		id entity.ArticleID,
		article *entity.Article,
	) error
	GetArticles(
		ctx context.Context,
		cursorPublishedAt int64,
		cursorTitle string,
		order CursorOrder,
		n int,
		articles *[]entity.Article,
	) error

	GetArticleHTML(
		ctx context.Context,
		id entity.ArticleID,
		body *[]byte,
	) error

	ConvertMD(
		ctx context.Context,
		source []byte,
		article *entity.Article,
		converted *[]byte,
	) error
	UpdateArticleByID(
		ctx context.Context,
		bucket string,
		articleID entity.ArticleID,
	) error

	GenerateBlogSiteMap(
		ctx context.Context,
		origin string,
	) (*XMLURLSet, error)

	DeleteArticle(
		ctx context.Context,
		articleID entity.ArticleID,
	) error
}

// Impl ...
type Impl struct {
	db          DB
	storage     Storage
	converterMD cmarkdown.Converter
}

// NewImpl ...
func NewImpl(
	db DB,
	storage Storage,
	converterMD cmarkdown.Converter,
) *Impl {
	return &Impl{
		db:          db,
		storage:     storage,
		converterMD: converterMD,
	}
}

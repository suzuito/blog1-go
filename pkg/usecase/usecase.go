package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/common-go/cmarkdown"
)

// Usecase ...
type Usecase interface {
	SyncArticles(
		ctx context.Context,
		source ArticleReader,
	) error

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

	ConvertMD(
		ctx context.Context,
		source []byte,
		article *entity.Article,
		converted *[]byte,
	) error

	GenerateBlogSiteMap(
		ctx context.Context,
		origin string,
	) (*XMLURLSet, error)

	GetAdminAuth(
		ctx context.Context,
		headerAdminAuth string,
		adminAuth *entity.AdminAuth,
	) error

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

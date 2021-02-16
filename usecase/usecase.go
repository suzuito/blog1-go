package usecase

import (
	"context"
	"net/http"

	"github.com/suzuito/blog1-go/entity/model"
	"github.com/suzuito/common-go/clogger"
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

	GenerateBlogSiteMap(
		ctx context.Context,
		origin string,
	) (string, error)

	GetAdminAuth(
		ctx context.Context,
		headerAdminAuth string,
		adminAuth *model.AdminAuth,
	) error

	ServeFront(
		ctx context.Context,
		w http.ResponseWriter,
		r *http.Request,
	) error
}

// Impl ...
type Impl struct {
	db          DB
	storage     Storage
	converterMD MDConverter
	logger      clogger.Logger
}

// NewImpl ...
func NewImpl(
	logger clogger.Logger,
	db DB,
	storage Storage,
	converterMD MDConverter,
) *Impl {
	return &Impl{
		logger:      logger,
		db:          db,
		storage:     storage,
		converterMD: converterMD,
	}
}

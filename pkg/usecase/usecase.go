package usecase

import (
	"context"

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

	GetArticleMarkdown(
		ctx context.Context,
		articleID entity.ArticleID,
		dst *[]byte,
	) error

	UpdateArticle(
		ctx context.Context,
		article *entity.Article,
		htmlString string,
	) error

	ConvertFromMarkdownToHTML(
		ctx context.Context,
		srcMD []byte,
		retHTML *string,
		article *entity.Article,
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
	MDConverter      MDConverter
	DB               DB
	Storage          Storage
	HTMLEditor       HTMLEditor
	HTMLMediaFetcher HTMLMediaFetcher
	HTMLTOCExtractor HTMLTOCExtractor
}

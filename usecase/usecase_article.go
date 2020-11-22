package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/entity/model"
	"golang.org/x/xerrors"
)

// GetArticles ...
func (u *Impl) GetArticles(
	ctx context.Context,
	startPublishedAt int64,
	n int,
	articles *[]model.Article,
) error {
	// return u.db.GetArticles(ctx, startPublishedAt, n, articles)
	return nil
}

// GetArticle ...
func (u *Impl) GetArticle(
	ctx context.Context,
	articleID model.ArticleID,
	article *model.Article,
) error {
	// return u.db.GetArticle(ctx, articleID, article)
	return nil
}

// CreateArticle ...
func (u *Impl) CreateArticle(
	ctx context.Context,
	article *model.Article,
) error {
	return u.CreateArticle(ctx, article)
}

// SyncArticles ...
func (u *Impl) SyncArticles(
	ctx context.Context,
	source ArticleReader,
) error {
	if err := source.Walk(ctx, func(article *model.Article) error {
		if err := u.db.SetArticle(ctx, article); err != nil {
			return xerrors.Errorf("Cannot set article '%+v' : %w", article, err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("ArticleReader walk is failed : %w", err)
	}
	return nil
}

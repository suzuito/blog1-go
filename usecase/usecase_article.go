package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/entity/model"
	"golang.org/x/xerrors"
)

// GetArticles ...
func (u *Impl) GetArticles(
	ctx context.Context,
	cursorPublishedAt int64,
	cursorTitle string,
	order CursorOrder,
	n int,
	articles *[]model.Article,
) error {
	return u.db.GetArticles(
		ctx,
		cursorPublishedAt,
		cursorTitle,
		order,
		n,
		articles,
	)
}

// GetArticle ...
func (u *Impl) GetArticle(
	ctx context.Context,
	articleID model.ArticleID,
	article *model.Article,
) error {
	return u.db.GetArticle(ctx, articleID, article)
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
	if err := source.Walk(ctx, func(article *model.Article, _ []byte) error {
		if err := u.db.SetArticle(ctx, article); err != nil {
			return xerrors.Errorf("Cannot set article '%+v' : %w", article, err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("ArticleReader walk is failed : %w", err)
	}
	return nil
}

// UploadArticleMDs ...
func (u *Impl) UploadArticleMDs(
	ctx context.Context,
	source ArticleReader,
) error {
	u.logger.Infof("Upload article MarkDown")
	if err := source.Walk(ctx, func(article *model.Article, raw []byte) error {
		converted := []byte{}
		if err := u.converterMD.Convert(ctx, article, raw, &converted); err != nil {
			return xerrors.Errorf("Cannot convert article '%+v' : %w", article, err)
		}
		u.logger.Infof("Upload '%s'", article.ID)
		if err := u.storage.UploadArticle(ctx, article, string(converted)); err != nil {
			return xerrors.Errorf("Cannot upload article '%+v' : %w", article, err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("ArticleReader walk is failed : %w", err)
	}
	return nil
}

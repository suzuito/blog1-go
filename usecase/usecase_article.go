package usecase

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

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
	if err := source.Walk(ctx, func(article *model.Article, raw []byte) error {
		converted := []byte{}
		if err := u.converterMD.Convert(ctx, article, raw, &converted); err != nil {
			return xerrors.Errorf("Cannot convert article '%+v' : %w", article, err)
		}
		u.logger.Infof("Upload '%s'", article.ID)
		if err := u.storage.UploadArticle(ctx, article, string(converted)); err != nil {
			return xerrors.Errorf("Cannot upload article '%+v' : %w", article, err)
		}
		if err := u.attacheArticleImages(article, converted); err != nil {
			return xerrors.Errorf("Cannot attacheArticleImages : %w", err)
		}
		if err := u.db.SetArticle(ctx, article); err != nil {
			return xerrors.Errorf("Cannot set article '%+v' : %w", article, err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("ArticleReader walk is failed : %w", err)
	}
	return nil
}

// WriteArticleHTMLs ...
func (u *Impl) WriteArticleHTMLs(
	ctx context.Context,
	source ArticleReader,
) error {
	if err := source.Walk(ctx, func(article *model.Article, raw []byte) error {
		converted := []byte{}
		if err := u.converterMD.Convert(ctx, article, raw, &converted); err != nil {
			return xerrors.Errorf("Cannot convert article '%+v' : %w", article, err)
		}
		filePath := fmt.Sprintf(".output/%s.html", article.Title)
		u.logger.Infof("Write '%s' into '%s'", article.ID, filePath)
		if err := os.RemoveAll(filePath); err != nil {
			return xerrors.Errorf("Cannot remove '%s'", filePath)
		}
		if err := ioutil.WriteFile(filePath, converted, 0644); err != nil {
			return xerrors.Errorf("Cannot write '%s'", filePath)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("ArticleReader walk is failed : %w", err)
	}
	return nil
}

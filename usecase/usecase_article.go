package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/entity/model"
)

// GetArticles ...
func (u *Usecase) GetArticles(
	ctx context.Context,
	startPublishedAt int64,
	n int,
	articles *[]model.Article,
) error {
	return u.db.GetArticles(ctx, startPublishedAt, n, articles)
}

// GetArticle ...
func (u *Usecase) GetArticle(
	ctx context.Context,
	articleID model.ArticleID,
	article *model.Article,
) error {
	return u.db.GetArticle(ctx, articleID, article)
}

// CreateArticle ...
func (u *Usecase) CreateArticle(
	ctx context.Context,
	article *model.Article,
) error {
	return u.CreateArticle(ctx, article)
}

package fdb

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/blog1-go/entity/model"
)

// GetArticle ...
func (c *ClientFirestore) GetArticle(
	ctx context.Context,
	articleID model.ArticleID,
	article *model.Article,
) error {
	return getDoc(
		ctx,
		c.cli.Collection(CollArticles),
		fmt.Sprintf("%s/%s", CollArticles, articleID),
		article,
	)
}

// GetArticles ...
func (c *ClientFirestore) GetArticles(
	ctx context.Context,
	articleID model.ArticleID,
	articles *[]model.Article,
) error {
	return getDocs(
		ctx,
		c.cli.Collection(CollArticles),
		func(shp *firestore.DocumentSnapshot) error {
			article := model.Article{}
			if err := shp.DataTo(&article); err != nil {
				return err
			}
			*articles = append(*articles, &article)
			return nil
		},
	)
}

// CreateArticle ...
func (c *ClientFirestore) CreateArticle(
	ctx context.Context,
	article *model.Article,
) error {
	return setDoc(
		ctx,
		c.cli.Collection(CollArticles),
		article,
	)
}

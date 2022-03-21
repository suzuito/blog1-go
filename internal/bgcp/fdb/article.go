package fdb

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/blog1-go/pkg/usecase"
	"golang.org/x/xerrors"
)

// GetArticle ...
func (c *Client) GetArticle(
	ctx context.Context,
	articleID entity.ArticleID,
	article *entity.Article,
) error {
	return getDoc(
		ctx,
		c.cli.Collection(CollArticles),
		string(articleID),
		article,
	)
}

// GetArticles ...
func (c *Client) GetArticles(
	ctx context.Context,
	cursorPublishedAt int64,
	cursorTitle string,
	order usecase.CursorOrder,
	n int,
	articles *[]entity.Article,
) error {
	coll := c.cli.Collection(CollArticles)
	query := coll.
		OrderBy("PublishedAt", newFirestoreOrder(order)).
		OrderBy("Title", firestore.Asc).
		StartAfter(cursorPublishedAt, cursorTitle).
		Limit(n)
	return getDocs(ctx, &query, func(shp *firestore.DocumentSnapshot) error {
		article := entity.Article{}
		if err := shp.DataTo(&article); err != nil {
			return errors.Wrapf(err, "Cannot DataTo from %s", coll.ID)
		}
		*articles = append(*articles, article)
		return nil
	})
}

// CreateArticle ...
func (c *Client) CreateArticle(
	ctx context.Context,
	article *entity.Article,
) error {
	now := time.Now()
	article.ID = entity.ArticleID(fmt.Sprintf("%d-%s", article.PublishedAt, article.Title))
	article.CreatedAt = now.Unix()
	article.UpdatedAt = now.Unix()
	if err := article.Validate(); err != nil {
		return errors.Wrapf(err, "Invalid article '%+v'", article)
	}
	return c.cli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		collArticles := c.cli.Collection(CollArticles)
		if err := getDocByTx(tx, collArticles, string(article.ID), &entity.Article{}); err != nil {
			if !xerrors.Is(err, usecase.ErrNotFound) {
				return errors.Wrapf(err, "Getting doc '%+v' is error", article)
			}
		} else {
			return errors.Wrapf(usecase.ErrAlreadyExists, "Doc '%+v' already exists error", article)
		}
		if err := setDocByTx(tx, collArticles, string(article.ID), article); err != nil {
			return errors.Wrapf(err, "Setting doc '%+v' is error", article)
		}
		collTags := c.cli.Collection(CollTags)
		for _, tag := range article.Tags {
			if err := setDocByTx(tx, collTags, string(tag.Name), tag); err != nil {
				return errors.Wrapf(err, "Setting doc '%+v' is error", tag)
			}
			collTagArticles := collTags.Doc(tag.Name).Collection(CollArticles)
			if err := setDocByTx(tx, collTagArticles, string(article.ID), article); err != nil {
				return errors.Wrapf(err, "Setting doc '%+v' is error", tag)
			}
		}
		return nil
	})
}

type diffTagType int

var (
	diffTagTypeCreated diffTagType = 1
	diffTagTypeDeleted diffTagType = 2
	diffTagTypeUpdated diffTagType = 3
)

type diffTag struct {
	Tag  entity.Tag
	Type diffTagType
}

func getDiffTags(
	beforeTags []entity.Tag,
	afterTags []entity.Tag,
) *[]diffTag {
	r := []diffTag{}
	for _, beforeTag := range beforeTags {
		exists := false
		for _, afterTag := range afterTags {
			if beforeTag.Name == afterTag.Name {
				exists = true
				break
			}
		}
		if !exists {
			r = append(r, diffTag{
				Tag:  beforeTag,
				Type: diffTagTypeDeleted,
			})
		}
	}
	for _, afterTag := range afterTags {
		exists := false
		for _, beforeTag := range beforeTags {
			if beforeTag.Name == afterTag.Name {
				exists = true
				break
			}
		}
		if !exists {
			r = append(r, diffTag{
				Tag:  afterTag,
				Type: diffTagTypeCreated,
			})
		}
	}
	for _, beforeTag := range beforeTags {
		for _, afterTag := range afterTags {
			if beforeTag.Name == afterTag.Name {
				r = append(r, diffTag{
					Tag:  beforeTag,
					Type: diffTagTypeUpdated,
				})
			}
		}
	}
	return &r
}

// SetArticle ...
func (c *Client) SetArticle(
	ctx context.Context,
	afterArticle *entity.Article,
) error {
	now := time.Now()
	if afterArticle.CreatedAt == 0 {
		afterArticle.CreatedAt = now.Unix()
	}
	afterArticle.UpdatedAt = now.Unix()
	if err := afterArticle.Validate(); err != nil {
		return errors.Wrapf(err, "Invalid article '%+v'", afterArticle)
	}
	return c.cli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		collArticles := c.cli.Collection(CollArticles)
		beforeArticle := entity.Article{}
		if err := getDocByTx(tx, collArticles, string(afterArticle.ID), &beforeArticle); err != nil {
			if !xerrors.Is(err, usecase.ErrNotFound) {
				return errors.Wrapf(err, "Get doc '%+v' is error", afterArticle)
			}
		}
		diffs := getDiffTags(beforeArticle.Tags, afterArticle.Tags)
		if err := setDocByTx(tx, collArticles, string(afterArticle.ID), afterArticle); err != nil {
			return errors.Wrapf(err, "Setting doc '%+v' is error", afterArticle)
		}
		collTags := c.cli.Collection(CollTags)
		for _, diff := range *diffs {
			collTagsArticles := collTags.Doc(diff.Tag.Name).Collection(CollArticles)
			if diff.Type == diffTagTypeCreated {
				if err := setDocByTx(tx, collTags, string(diff.Tag.Name), diff.Tag); err != nil {
					return errors.Wrapf(err, "Setting doc '%+v' is error", diff.Tag)
				}
				if err := setDocByTx(tx, collTagsArticles, string(afterArticle.ID), afterArticle); err != nil {
					return errors.Wrapf(err, "Setting doc '%+v' is error", afterArticle)
				}
			}
			if diff.Type == diffTagTypeDeleted {
				if err := deleteDocByTx(tx, collTags, string(diff.Tag.Name)); err != nil {
					return errors.Wrapf(err, "Set doc '%+v' is error", diff.Tag)
				}
				if err := deleteDocByTx(tx, collTagsArticles, string(afterArticle.ID)); err != nil {
					return errors.Wrapf(err, "Set doc '%+v' is error", afterArticle)
				}
			}
			if diff.Type == diffTagTypeUpdated {
				if err := setDocByTx(tx, collTagsArticles, string(afterArticle.ID), afterArticle); err != nil {
					return errors.Wrapf(err, "Set doc '%+v' is error", afterArticle)
				}
			}
		}
		return nil
	})
}

// DeleteArticle ...
func (c *Client) DeleteArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) error {
	return c.cli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		collArticles := c.cli.Collection(CollArticles)
		if err := deleteDocByTx(tx, collArticles, string(articleID)); err != nil {
			return errors.Wrapf(err, "cannot delete article %s", articleID)
		}
		return nil
	})
}

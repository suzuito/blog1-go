package fdb

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/blog1-go/entity/model"
	"github.com/suzuito/blog1-go/usecase"
	"golang.org/x/xerrors"
)

// GetArticle ...
func (c *Client) GetArticle(
	ctx context.Context,
	articleID model.ArticleID,
	article *model.Article,
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
	articles *[]model.Article,
) error {
	coll := c.cli.Collection(CollArticles)
	query := coll.
		OrderBy("PublishedAt", newFirestoreOrder(order)).
		OrderBy("Title", firestore.Asc).
		StartAfter(cursorPublishedAt, cursorTitle).
		Limit(n)
	return getDocs(ctx, &query, func(shp *firestore.DocumentSnapshot) error {
		article := model.Article{}
		if err := shp.DataTo(&article); err != nil {
			return xerrors.Errorf("Cannot data to : %w", err)
		}
		*articles = append(*articles, article)
		return nil
	})
}

// CreateArticle ...
func (c *Client) CreateArticle(
	ctx context.Context,
	article *model.Article,
) error {
	now := time.Now()
	article.ID = model.ArticleID(fmt.Sprintf("%d-%s", article.PublishedAt, article.Title))
	article.CreatedAt = now.Unix()
	article.UpdatedAt = now.Unix()
	if err := article.Validate(); err != nil {
		return xerrors.Errorf("Invalid article '%+v' : %w", article, err)
	}
	return c.cli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		collArticles := c.cli.Collection(CollArticles)
		if err := getDocByTx(tx, collArticles, string(article.ID), &model.Article{}); err != nil {
			if !xerrors.Is(err, usecase.ErrNotFound) {
				return xerrors.Errorf("Get doc '%+v' is error : %w", article, err)
			}
		} else {
			return xerrors.Errorf("Doc '%+v' already exists error : %w", article, usecase.ErrAlreadyExists)
		}
		if err := setDocByTx(tx, collArticles, string(article.ID), article); err != nil {
			return xerrors.Errorf("Set doc '%+v' is error : %w", article, err)
		}
		collTags := c.cli.Collection(CollTags)
		for _, tag := range article.Tags {
			if err := setDocByTx(tx, collTags, string(tag.Name), tag); err != nil {
				return xerrors.Errorf("Set doc '%+v' is error : %w", tag, err)
			}
			collTagArticles := collTags.Doc(tag.Name).Collection(CollArticles)
			if err := setDocByTx(tx, collTagArticles, string(article.ID), article); err != nil {
				return xerrors.Errorf("Set doc '%+v' is error : %w", tag, err)
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
	Tag  model.Tag
	Type diffTagType
}

func getDiffTags(
	beforeTags []model.Tag,
	afterTags []model.Tag,
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
	afterArticle *model.Article,
) error {
	now := time.Now()
	if afterArticle.CreatedAt == 0 {
		afterArticle.CreatedAt = now.Unix()
	}
	afterArticle.UpdatedAt = now.Unix()
	if err := afterArticle.Validate(); err != nil {
		return xerrors.Errorf("Invalid article '%+v' : %w", afterArticle, err)
	}
	return c.cli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		collArticles := c.cli.Collection(CollArticles)
		beforeArticle := model.Article{}
		if err := getDocByTx(tx, collArticles, string(afterArticle.ID), &beforeArticle); err != nil {
			if !xerrors.Is(err, usecase.ErrNotFound) {
				return xerrors.Errorf("Get doc '%+v' is error : %w", afterArticle, err)
			}
		}
		diffs := getDiffTags(beforeArticle.Tags, afterArticle.Tags)
		if err := setDocByTx(tx, collArticles, string(afterArticle.ID), afterArticle); err != nil {
			return xerrors.Errorf("Set doc '%+v' is error : %w", afterArticle, err)
		}
		collTags := c.cli.Collection(CollTags)
		for _, diff := range *diffs {
			collTagsArticles := collTags.Doc(diff.Tag.Name).Collection(CollArticles)
			if diff.Type == diffTagTypeCreated {
				if err := setDocByTx(tx, collTags, string(diff.Tag.Name), diff.Tag); err != nil {
					return xerrors.Errorf("Set doc '%+v' is error : %w", diff.Tag, err)
				}
				if err := setDocByTx(tx, collTagsArticles, string(afterArticle.ID), afterArticle); err != nil {
					return xerrors.Errorf("Set doc '%+v' is error : %w", afterArticle, err)
				}
			}
			if diff.Type == diffTagTypeDeleted {
				if err := deleteDocByTx(tx, collTags, string(diff.Tag.Name)); err != nil {
					return xerrors.Errorf("Set doc '%+v' is error : %w", diff.Tag, err)
				}
				if err := deleteDocByTx(tx, collTagsArticles, string(afterArticle.ID)); err != nil {
					return xerrors.Errorf("Set doc '%+v' is error : %w", afterArticle, err)
				}
			}
			if diff.Type == diffTagTypeUpdated {
				if err := setDocByTx(tx, collTagsArticles, string(afterArticle.ID), afterArticle); err != nil {
					return xerrors.Errorf("Set doc '%+v' is error : %w", afterArticle, err)
				}
			}
		}
		return nil
	})
}

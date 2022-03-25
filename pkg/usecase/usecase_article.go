package usecase

import (
	"bytes"
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/internal/cmarkdown"
	"github.com/suzuito/blog1-go/pkg/entity"
)

// GetArticles ...
func (u *Impl) GetArticles(
	ctx context.Context,
	cursorPublishedAt int64,
	cursorTitle string,
	order CursorOrder,
	n int,
	articles *[]entity.Article,
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
	articleID entity.ArticleID,
	article *entity.Article,
) error {
	return u.db.GetArticle(ctx, articleID, article)
}

// CreateArticle ...
func (u *Impl) CreateArticle(
	ctx context.Context,
	article *entity.Article,
) error {
	return u.CreateArticle(ctx, article)
}

func (u *Impl) UpdateArticleByID(
	ctx context.Context,
	bucket string,
	articleID entity.ArticleID,
) error {
	path := fmt.Sprintf("%s/%s.md", bucket, articleID)
	return u.UpdateArticle(ctx, path)
}

func (u *Impl) UpdateArticle(
	ctx context.Context,
	path string,
) error {
	source := []byte{}
	headers := map[string]string{}
	if err := u.storage.GetFileAsHTTPResponse(ctx, path, &source, &headers); err != nil {
		return errors.Wrapf(err, "cannot get file from %s", path)
	}
	converted := []byte{}
	output := bytes.NewBufferString("")
	tocs := []cmarkdown.CMTOC{}
	images := []cmarkdown.CMImage{}
	meta := cmarkdown.CMMeta{}
	if err := u.converterMD.Convert(ctx, source, output, &meta, &tocs, &images); err != nil {
		return errors.Wrapf(err, "cannot convert")
	}
	article := newArticleFromCMeta(&meta, tocs, images)
	converted = output.Bytes()
	// FIXME into usecase
	// u.refineArticleImage(http.DefaultClient, &articleImage)
	if article.ID == entity.ArticleID("") {
		return errors.Errorf("Empty ID is detected '%+v'", article)
	}
	fmt.Printf("Upload '%s'\n", article.ID)
	if err := u.storage.UploadArticle(ctx, article, string(converted)); err != nil {
		return errors.Wrapf(err, "cannot upload article '%+v'", article)
	}
	if err := u.attacheArticleImages(article, converted); err != nil {
		return errors.Wrapf(err, "cannot attacheArticleImages")
	}
	if err := u.db.SetArticle(ctx, article); err != nil {
		return errors.Wrapf(err, "cannot set article '%+v'", article)
	}
	return nil
}

func (u *Impl) DeleteArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) error {
	fmt.Printf("Delete '%s'\n", articleID)
	if err := u.storage.DeleteArticle(ctx, articleID); err != nil {
		return errors.Wrapf(err, "cannot delete article '%s'", articleID)
	}
	if err := u.db.DeleteArticle(ctx, articleID); err != nil {
		return errors.Wrapf(err, "cannot delete article '%s'", articleID)
	}
	return nil
}

func (u *Impl) ConvertMD(
	ctx context.Context,
	source []byte,
	article *entity.Article,
	converted *[]byte,
) error {
	output := bytes.NewBufferString("")
	tocs := []cmarkdown.CMTOC{}
	images := []cmarkdown.CMImage{}
	meta := cmarkdown.CMMeta{}
	if err := u.converterMD.Convert(ctx, source, output, &meta, &tocs, &images); err != nil {
		return errors.Wrapf(err, "cannot convert")
	}
	*article = *newArticleFromCMeta(&meta, tocs, images)
	*converted = output.Bytes()
	return nil
}

func (u *Impl) GetArticleHTML(
	ctx context.Context,
	id entity.ArticleID,
	body *[]byte,
) error {
	path := fmt.Sprintf("%s.html", id)
	return u.storage.GetFileAsHTTPResponse(
		ctx,
		path,
		body,
		&map[string]string{},
	)
}

func newArticleFromCMeta(
	meta *cmarkdown.CMMeta,
	tocs []cmarkdown.CMTOC,
	images []cmarkdown.CMImage,
) *entity.Article {
	a := entity.Article{}
	a.ID = entity.ArticleID(meta.ID)
	a.Description = meta.Description
	a.PublishedAt = meta.DateAsTime().Unix()
	a.Tags = *entity.NewTags(meta.Tags)
	a.Title = meta.Title
	for _, toc := range tocs {
		articleIndex := entity.ArticleIndex{
			ID:    toc.ID,
			Name:  toc.Name,
			Level: entity.ArticleIndexLevel(toc.Level),
		}
		a.TOC = append(a.TOC, articleIndex)
	}
	for _, image := range images {
		articleImage := entity.ArticleImage{
			URL: image.URL,
		}
		a.Images = append(a.Images, articleImage)
	}
	return &a
}

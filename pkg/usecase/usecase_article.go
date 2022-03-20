package usecase

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/common-go/cmarkdown"
	"golang.org/x/xerrors"
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
		return xerrors.Errorf("Cannot get file from %s : %w", path, err)
	}
	converted := []byte{}
	article := entity.Article{}
	if err := u.ConvertMD(ctx, source, &article, &converted); err != nil {
		return xerrors.Errorf("Cannot convert article '%+v' : %w", article, err)
	}
	if article.ID == entity.ArticleID("") {
		return xerrors.Errorf("Empty ID is detected '%+v'", article)
	}
	fmt.Printf("Upload '%s'\n", article.ID)
	if err := u.storage.UploadArticle(ctx, &article, string(converted)); err != nil {
		return xerrors.Errorf("Cannot upload article '%+v' : %w", article, err)
	}
	if err := u.attacheArticleImages(&article, converted); err != nil {
		return xerrors.Errorf("Cannot attacheArticleImages : %w", err)
	}
	if err := u.db.SetArticle(ctx, &article); err != nil {
		return xerrors.Errorf("Cannot set article '%+v' : %w", article, err)
	}
	return nil
}

func (u *Impl) DeleteArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) error {
	fmt.Printf("Delete '%s'\n", articleID)
	if err := u.storage.DeleteArticle(ctx, articleID); err != nil {
		return xerrors.Errorf("Cannot delete article '%s' : %w", articleID, err)
	}
	if err := u.db.DeleteArticle(ctx, articleID); err != nil {
		return xerrors.Errorf("Cannot delete article '%s' : %w", articleID, err)
	}
	return nil
}

// SyncArticles ...
// :Deprecated
func (u *Impl) SyncArticles(
	ctx context.Context,
	source ArticleReader,
) error {
	if err := source.Walk(ctx, func(article *entity.Article, raw []byte) error {
		converted := []byte{}
		if err := u.ConvertMD(ctx, raw, article, &converted); err != nil {
			return xerrors.Errorf("Cannot convert article '%+v' : %w", article, err)
		}
		fmt.Printf("Upload '%s'\n", article.ID)
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
// :Deprecated
func (u *Impl) WriteArticleHTMLs(
	ctx context.Context,
	source ArticleReader,
) error {
	if err := source.Walk(ctx, func(article *entity.Article, raw []byte) error {
		converted := []byte{}
		if err := u.ConvertMD(ctx, raw, article, &converted); err != nil {
			return xerrors.Errorf("Cannot convert article '%+v' : %w", article, err)
		}
		filePath := fmt.Sprintf(".output/%s.html", article.Title)
		fmt.Printf("Write '%s' into '%s'", article.ID, filePath)
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
		return xerrors.Errorf("Cannot convert : %w", err)
	}
	article.ID = entity.ArticleID(meta.ID)
	article.Description = meta.Description
	article.PublishedAt = meta.DateAsTime().Unix()
	article.Tags = *entity.NewTags(meta.Tags)
	article.Title = meta.Title
	for _, toc := range tocs {
		articleIndex := entity.ArticleIndex{
			ID:    toc.ID,
			Name:  toc.Name,
			Level: entity.ArticleIndexLevel(toc.Level),
		}
		article.TOC = append(article.TOC, articleIndex)
	}
	for _, image := range images {
		articleImage := entity.ArticleImage{
			URL: image.URL,
		}
		u.refineArticleImage(http.DefaultClient, &articleImage)
		article.Images = append(article.Images, articleImage)
	}
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

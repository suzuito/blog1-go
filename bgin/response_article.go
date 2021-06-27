package bgin

import (
	"github.com/suzuito/blog1-go/entity/model"
)

// ResponseArticle ...
type ResponseArticle struct {
	ID          model.ArticleID        `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	CreatedAt   int64                  `json:"createdAt"`
	UpdatedAt   int64                  `json:"updatedAt"`
	PublishedAt int64                  `json:"publishedAt"`
	Tags        []ResponseTag          `json:"tags"`
	Images      []ResponseArticleImage `json:"images"`
	TOC         []ResponseArticleIndex `json:"toc"`
}

// NewResponseArticle ...
func NewResponseArticle(a *model.Article) *ResponseArticle {
	imgs := []ResponseArticleImage{}
	for _, aimg := range a.Images {
		imgs = append(imgs, *NewResponseArticleImage(&aimg))
	}
	indexes := []ResponseArticleIndex{}
	for _, index := range a.TOC {
		indexes = append(indexes, ResponseArticleIndex{
			ID:    index.ID,
			Name:  index.Name,
			Level: index.Level,
		})
	}
	return &ResponseArticle{
		ID:          a.ID,
		Description: a.Description,
		Title:       a.Title,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
		PublishedAt: a.PublishedAt,
		Tags:        *NewResponseTags(&a.Tags),
		Images:      imgs,
		TOC:         indexes,
	}
}

// NewResponseArticles ...
func NewResponseArticles(a *[]model.Article) *[]ResponseArticle {
	b := []ResponseArticle{}
	for _, v := range *a {
		b = append(b, *NewResponseArticle(&v))
	}
	return &b
}

// ResponseArticleImage ...
type ResponseArticleImage struct {
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	URL        string `json:"url"`
	RealWidth  int    `json:"realWidth"`
	RealHeight int    `json:"realHeight"`
}

// NewResponseArticleImage ...
func NewResponseArticleImage(a *model.ArticleImage) *ResponseArticleImage {
	return &ResponseArticleImage{
		Width:      a.Width,
		Height:     a.Height,
		URL:        a.URL,
		RealWidth:  a.RealWidth,
		RealHeight: a.RealHeight,
	}
}

type ResponseArticleIndex struct {
	ID    string                  `json:"id"`
	Name  string                  `json:"name"`
	Level model.ArticleIndexLevel `json:"level"`
}

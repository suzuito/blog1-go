package bgin

import "github.com/suzuito/blog1-go/entity/model"

// ResponseArticle ...
type ResponseArticle struct {
	ID          model.ArticleID         `json:"id"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	CreatedAt   int64                   `json:"createdAt"`
	UpdatedAt   int64                   `json:"updatedAt"`
	PublishedAt int64                   `json:"publishedAt"`
	Versions    ResponseArticleVersions `json:"versions"`
	Tags        []ResponseTag           `json:"tags"`
}

// NewResponseArticle ...
func NewResponseArticle(a *model.Article) *ResponseArticle {
	return &ResponseArticle{
		ID:          a.ID,
		Description: a.Description,
		Title:       a.Title,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
		PublishedAt: a.PublishedAt,
		Versions:    *NewResponseArticleVersions(&a.Versions),
		Tags:        *NewResponseTags(&a.Tags),
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

// ResponseArticleVersions ...
type ResponseArticleVersions struct {
	Current int64 `json:"current"`
}

// NewResponseArticleVersions ...
func NewResponseArticleVersions(a *model.ArticleVersions) *ResponseArticleVersions {
	return &ResponseArticleVersions{
		Current: a.Current,
	}
}

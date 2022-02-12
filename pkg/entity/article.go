package entity

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	// ErrArticleMetaBlockNotFound ...
	ErrArticleMetaBlockNotFound = fmt.Errorf("Article's meta block is not found")
)

// ArticleID ...
type ArticleID string

// Article ...
type Article struct {
	ID          ArticleID `validate:"required"`
	Title       string    `validate:"required"`
	Description string
	CreatedAt   int64
	UpdatedAt   int64
	PublishedAt int64
	Tags        []Tag
	Images      []ArticleImage
	TOC         []ArticleIndex
}

// Validate ...
func (a *Article) Validate() error {
	validator := validator.New()
	return validator.Struct(a)
}

// ArticleIndexLevel ...
type ArticleIndexLevel int

const (
	ArticleIndexLevel1 ArticleIndexLevel = 1
	ArticleIndexLevel2 ArticleIndexLevel = 2
	ArticleIndexLevel3 ArticleIndexLevel = 3
	ArticleIndexLevel4 ArticleIndexLevel = 4
	ArticleIndexLevel5 ArticleIndexLevel = 5
)

// ArticleIndex ...
type ArticleIndex struct {
	ID    string
	Name  string
	Level ArticleIndexLevel
}

// ArticleImage ...
type ArticleImage struct {
	Width      int
	Height     int
	URL        string
	RealWidth  int
	RealHeight int
}

// Tag ...
type Tag struct {
	Name string
}

// NewTags ...
func NewTags(tags []string) *[]Tag {
	r := []Tag{}
	for _, tag := range tags {
		r = append(r, Tag{Name: tag})
	}
	return &r
}

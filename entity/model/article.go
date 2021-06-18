package model

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/xerrors"
	"gopkg.in/yaml.v2"
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

// ArticleIndex ...
type ArticleIndex struct {
	Name     string
	Children []ArticleIndex
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

// NewArticleFromRawContent ...
func NewArticleFromRawContent(r io.Reader) (*Article, []byte, error) {
	s := bufio.NewScanner(r)
	isMetaBlock := false
	isMetaBlockDone := false
	metaBlock := ""
	notMetaBlock := ""
	for s.Scan() {
		l := s.Text()
		if strings.HasPrefix(l, "---") && !isMetaBlockDone {
			if !isMetaBlock {
				isMetaBlock = true
				continue
			}
			isMetaBlock = false
			isMetaBlockDone = true
			continue
		}
		if isMetaBlock {
			metaBlock += l + "\n"
		} else {
			notMetaBlock += l + "\n"
		}
	}
	if !isMetaBlockDone {
		return nil, nil, xerrors.Errorf("Meta data is not found : %w", ErrArticleMetaBlockNotFound)
	}
	embedMeta := struct {
		Title       string   `yaml:"title"`
		Tags        []string `yaml:"tags"`
		Description string   `yaml:"description"`
		Date        string   `yaml:"date"`
	}{}
	if err := yaml.Unmarshal([]byte(metaBlock), &embedMeta); err != nil {
		return nil, nil, xerrors.Errorf("Cannot parse yaml block '%s' : %w", metaBlock, err)
	}
	date, err := time.Parse("2006-01-02", embedMeta.Date)
	if err != nil {
		return nil, nil, xerrors.Errorf("Cannot parse date '%s' : %w", embedMeta.Date, err)
	}
	article := Article{
		ID:          ArticleID(fmt.Sprintf("%s-%s", embedMeta.Date, embedMeta.Title)),
		Title:       embedMeta.Title,
		Description: embedMeta.Description,
		Tags:        *NewTags(embedMeta.Tags),
		PublishedAt: date.Unix(),
	}
	return &article, []byte(notMetaBlock), nil
}

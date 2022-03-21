package local

import (
	"context"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/entity"
)

// ArticleReaderFix ...
type ArticleReaderFix struct {
	filePaths []string
}

// NewArticleReaderFix ...
func NewArticleReaderFix() *ArticleReaderFix {
	return &ArticleReaderFix{
		filePaths: []string{},
	}
}

// AddFilePath ...
func (r *ArticleReaderFix) AddFilePath(filePath string) {
	r.filePaths = append(r.filePaths, filePath)
}

// Walk ...
func (r *ArticleReaderFix) Walk(ctx context.Context, each func(article *entity.Article, raw []byte) error) error {
	for _, filePath := range r.filePaths {
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			return errors.Wrapf(err, "reading file '%s' is failed", filePath)
		}
		article := entity.Article{}
		if err := each(&article, file); err != nil {
			return errors.Wrapf(err, "each is failed")
		}
	}
	return nil
}

// Close ...
func (r *ArticleReaderFix) Close() error {
	return nil
}

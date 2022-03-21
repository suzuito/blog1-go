package local

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/entity"
)

// ArticleReaderAll ...
type ArticleReaderAll struct {
	dirBase string
}

// NewArticleReaderAll ...
func NewArticleReaderAll(
	dirBase string,
) *ArticleReaderAll {
	return &ArticleReaderAll{
		dirBase: dirBase,
	}
}

// Walk ...
func (r *ArticleReaderAll) Walk(ctx context.Context, each func(article *entity.Article, raw []byte) error) error {
	err := filepath.Walk(r.dirBase, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.Wrapf(err, "Reading file '%s' is failed", path)
		}
		article := entity.Article{}
		return each(&article, file)
	})
	if err != nil {
		return errors.Wrapf(err, "Walk dir '%s' is failed", r.dirBase)
	}
	return nil
}

// Close ...
func (r *ArticleReaderAll) Close() error {
	return nil
}

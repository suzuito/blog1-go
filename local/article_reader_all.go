package local

import (
	"context"
	"os"
	"path/filepath"

	"github.com/suzuito/blog1-go/entity/model"
	"golang.org/x/xerrors"
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
func (r *ArticleReaderAll) Walk(ctx context.Context, each func(article *model.Article, raw []byte) error) error {
	err := filepath.Walk(r.dirBase, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return xerrors.Errorf("Reading file '%s' is failed : %w", path, err)
		}
		defer file.Close()
		article, raw, err := model.NewArticleFromRawContent(file)
		if err != nil {
			return xerrors.Errorf("Parse file '%s' is failed : %w", path, err)
		}
		return each(article, raw)
	})
	if err != nil {
		return xerrors.Errorf("Walk dir '%s' is failed : %w", r.dirBase, err)
	}
	return nil
}

// Close ...
func (r *ArticleReaderAll) Close() error {
	return nil
}

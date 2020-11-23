package local

import (
	"context"
	"os"
	"path/filepath"

	"github.com/suzuito/blog1-go/entity/model"
	"golang.org/x/xerrors"
)

// ArticleReader ...
type ArticleReader struct {
	dirBase string
}

// NewArticleReader ...
func NewArticleReader(
	dirBase string,
) *ArticleReader {
	return &ArticleReader{
		dirBase: dirBase,
	}
}

// Walk ...
func (r *ArticleReader) Walk(ctx context.Context, each func(article *model.Article, raw []byte) error) error {
	err := filepath.Walk(r.dirBase, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return xerrors.Errorf("Reading file '%s' is failed : %w", path, err)
		}
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
func (r *ArticleReader) Close() error {
	return nil
}

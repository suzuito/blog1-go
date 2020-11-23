package local

import (
	"bufio"
	"context"
	"os"

	"github.com/suzuito/blog1-go/entity/model"
	"golang.org/x/xerrors"
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
func (r *ArticleReaderFix) Walk(ctx context.Context, each func(article *model.Article, raw []byte) error) error {
	for _, filePath := range r.filePaths {
		fp, err := os.Open(filePath)
		if err != nil {
			return xerrors.Errorf("Cannot open file '%s' : %w", filePath, err)
		}
		file := bufio.NewReader(fp)
		article, raw, err := model.NewArticleFromRawContent(file)
		if err != nil {
			fp.Close()
			return xerrors.Errorf("Parse file '%s' is failed : %w", filePath, err)
		}
		if err := fp.Close(); err != nil {
			return xerrors.Errorf("Cannot close file '%s' : %w", filePath, err)
		}
		if err := each(article, raw); err != nil {
			return xerrors.Errorf("Error : %w", err)
		}
	}
	return nil
}

// Close ...
func (r *ArticleReaderFix) Close() error {
	return nil
}

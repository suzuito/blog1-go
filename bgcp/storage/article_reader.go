package storage

import (
	"context"
	"io/ioutil"

	gstorage "cloud.google.com/go/storage"
	"github.com/suzuito/blog1-go/entity/model"
	"golang.org/x/xerrors"
)

type ArticleReader struct {
	bucket string
	cli    *gstorage.Client
	key    string
}

func NewArticleReader(
	cli *gstorage.Client,
	bucket, key string,
) *ArticleReader {
	return &ArticleReader{
		cli: cli,
	}
}

// Walk ...
func (r *ArticleReader) Walk(ctx context.Context, each func(article *model.Article, raw []byte) error) error {
	bh := r.cli.Bucket(r.bucket)
	oh := bh.Object(r.key)
	reader, err := oh.NewReader(ctx)
	if err != nil {
		return xerrors.Errorf("Reading file '%s/%s' is failed : %w", r.bucket, r.key, err)
	}
	defer reader.Close()
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return xerrors.Errorf("Reading file '%s/%s' is failed : %w", r.bucket, r.key, err)
	}
	article := model.Article{}
	if err := each(&article, body); err != nil {
		return xerrors.Errorf("Error : %w", err)
	}
	return nil
}

// Close ...
func (r *ArticleReader) Close() error {
	return nil
}

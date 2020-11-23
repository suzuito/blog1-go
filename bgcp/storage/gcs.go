package storage

import (
	"context"
	"fmt"
	"io"
	"strings"

	gstorage "cloud.google.com/go/storage"
	"github.com/suzuito/blog1-go/entity/model"
	"golang.org/x/xerrors"
)

// GCS ...
type GCS struct {
	cli    *gstorage.Client
	bucket string
}

// New ...
func New(cli *gstorage.Client, bucket string) *GCS {
	return &GCS{cli: cli, bucket: bucket}
}

// UploadArticle ...
func (c *GCS) UploadArticle(
	ctx context.Context,
	article *model.Article,
	raw string,
) error {
	b := c.cli.Bucket(c.bucket)
	p := fmt.Sprintf("articles/%s.html", article.ID)
	o := b.Object(p)
	w := o.NewWriter(ctx)
	w.ContentType = "text/html;charset=utf-8"
	buf := strings.NewReader(raw)
	if _, err := io.Copy(w, buf); err != nil {
		return xerrors.Errorf("Cannot upload article '%s' into '%s/%s' : %w", article.ID, c.bucket, p, err)
	}
	if err := w.Close(); err != nil {
		return xerrors.Errorf("Cannot upload article '%s' into '%s/%s' : %w", article.ID, c.bucket, p, err)
	}
	return nil
}

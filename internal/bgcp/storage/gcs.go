package storage

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"cloud.google.com/go/storage"
	gstorage "cloud.google.com/go/storage"
	"github.com/suzuito/blog1-go/internal/setting"
	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/blog1-go/pkg/usecase"
	"golang.org/x/xerrors"
)

// NewResource ...
func NewResource(ctx context.Context, env *setting.Environment) (*gstorage.Client, error) {
	cli, err := gstorage.NewClient(ctx)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return cli, nil
}

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
	article *entity.Article,
	raw string,
) error {
	b := c.cli.Bucket(c.bucket)
	p := fmt.Sprintf("%s.html", article.ID)
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

func (c *GCS) GetFileAsHTTPResponse(
	ctx context.Context,
	p string,
	body *[]byte,
	headers *map[string]string,
) error {
	b := c.cli.Bucket(c.bucket)
	o := b.Object(p)
	reader, err := o.NewReader(ctx)
	if err != nil {
		if xerrors.Is(err, storage.ErrObjectNotExist) {
			return xerrors.Errorf("Not found '%s': %w", p, usecase.ErrNotFound)
		}
		return xerrors.Errorf("Cannot new reader '%s': %w", p, err)
	}
	defer reader.Close()
	*body, err = ioutil.ReadAll(reader)
	if err != nil {
		return xerrors.Errorf("Cannot read '%s': %w", p, err)
	}
	(*headers)["Content-Type"] = fmt.Sprintf("%s", reader.ContentType())
	return nil
}

func (c *GCS) UploadHTML(
	ctx context.Context,
	p string,
	body string,
) error {
	b := c.cli.Bucket(c.bucket)
	o := b.Object(p)
	w := o.NewWriter(ctx)
	w.ContentType = "text/html;charset=utf-8"
	buf := strings.NewReader(body)
	if _, err := io.Copy(w, buf); err != nil {
		return xerrors.Errorf("Cannot upload html into '%s/%s' : %w", c.bucket, p, err)
	}
	if err := w.Close(); err != nil {
		return xerrors.Errorf("Cannot upload html into '%s/%s' : %w", c.bucket, p, err)
	}
	return nil
}

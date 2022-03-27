package storage

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"cloud.google.com/go/storage"
	gstorage "cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/blog1-go/pkg/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
	"golang.org/x/xerrors"
)

// NewResource ...
func NewResource(ctx context.Context, env *setting.Environment) (*gstorage.Client, error) {
	cli, err := gstorage.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "storage.NewClient")
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
		return errors.Wrapf(err, "cannot upload article '%s' into '%s/%s'", article.ID, c.bucket, p)
	}
	if err := w.Close(); err != nil {
		return errors.Wrapf(err, "cannot upload article '%s' into '%s/%s'", article.ID, c.bucket, p)
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
			return errors.Wrapf(usecase.ErrNotFound, "not found '%s'", p)
		}
		return errors.Wrapf(err, "cannot new reader '%s'", p)
	}
	defer reader.Close()
	*body, err = ioutil.ReadAll(reader)
	if err != nil {
		return errors.Wrapf(err, "cannot read '%s'", p)
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
		return errors.Wrapf(err, "cannot upload html into '%s/%s'", c.bucket, p)
	}
	if err := w.Close(); err != nil {
		return errors.Wrapf(err, "Cannot upload html into '%s/%s'", c.bucket, p)
	}
	return nil
}

func (c *GCS) DeleteArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) error {
	b := c.cli.Bucket(c.bucket)
	p := fmt.Sprintf("%s.html", articleID)
	o := b.Object(p)
	if err := o.Delete(ctx); err != nil {
		if xerrors.Is(err, storage.ErrObjectNotExist) {
			return nil
		}
		return errors.Wrapf(err, "cannot delete from '%s/%s'", c.bucket, p)
	}
	return nil
}

var extractArticleIDFromPathRegexp = regexp.MustCompile(".md$")

func ExtractArticleIDFromPath(p string) entity.ArticleID {
	return entity.ArticleID(extractArticleIDFromPathRegexp.ReplaceAll([]byte(filepath.Base(p)), []byte("")))
}

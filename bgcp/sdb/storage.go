package sdb

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/suzuito/blog1-go/entity/model"
)

// ClientStorage ...
type ClientStorage struct {
	cli        *storage.Client
	bucketName string
	projectID  string
}

func (c *ClientStorage) bucket() *storage.BucketHandle {
	return c.cli.Bucket(c.bucketName)
}

// GetGenerations ...
func (c *ClientStorage) GetGenerations(ctx context.Context, articleID model.ArticleID, generations *[]int64) error {
	b := c.bucket()
	o := b.Object(articleID)
}

// SetDoc ...
func (c *ClientStorage) SetDoc(ctx context.Context, articleID model.ArticleID, body string) error {
	b := c.bucket()
	o := b.Object(articleID)
	w := o.NewWriter(ctx)
	if _, err := w.Write([]byte{body}); err != nil {
		return err
	}
	defer w.Close()
	return err
}

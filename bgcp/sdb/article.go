package sdb

import (
	"context"
)

// SetArticle ...
func (c *ClientStorage) SetArticle(
	ctx context.Context,
	body string,
) error {
	c.cli.Bucket()
}

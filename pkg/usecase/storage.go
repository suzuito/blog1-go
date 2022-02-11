package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/pkg/entity"
)

// Storage ...
type Storage interface {
	UploadArticle(
		ctx context.Context,
		article *entity.Article,
		raw string,
	) error
	GetFileAsHTTPResponse(
		ctx context.Context,
		p string,
		body *[]byte,
		headers *map[string]string,
	) error
	UploadHTML(
		ctx context.Context,
		p string,
		body string,
	) error
}

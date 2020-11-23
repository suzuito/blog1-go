package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/entity/model"
)

// Storage ...
type Storage interface {
	UploadArticle(
		ctx context.Context,
		article *model.Article,
		raw string,
	) error
}

package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/entity/model"
)

// MDConverter ...
type MDConverter interface {
	Convert(
		ctx context.Context,
		article *model.Article,
		raw []byte,
		converted *[]byte,
	) error
}

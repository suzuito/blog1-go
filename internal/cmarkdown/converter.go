package cmarkdown

import (
	"context"
)

type Converter interface {
	Convert(
		ctx context.Context,
		src string,
		dst *string,
		meta *CMMeta,
	) error
}

package cmarkdown

import (
	"context"
	"io"
)

type Converter interface {
	Convert(
		ctx context.Context,
		source []byte,
		w io.Writer,
		meta *CMMeta,
		tocs *[]CMTOC,
		images *[]CMImage,
	) error
}

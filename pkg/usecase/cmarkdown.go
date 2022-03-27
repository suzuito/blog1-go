package usecase

import (
	"context"
	"fmt"
	"time"
)

var ErrMetaNotFound = fmt.Errorf("Meta not found")

type CMMeta struct {
	ID          string   `yaml:"id"`
	Title       string   `yaml:"title"`
	Tags        []string `yaml:"tags"`
	Description string   `yaml:"description"`
	Date        string   `yaml:"date"`
}

func (c *CMMeta) DateAsTime() time.Time {
	r, err := time.Parse("2006-01-02", c.Date)
	if err != nil {
		return time.Unix(0, 0)
	}
	return r
}

type MDConverter interface {
	Convert(
		ctx context.Context,
		src string,
		dst *string,
		meta *CMMeta,
	) error
}

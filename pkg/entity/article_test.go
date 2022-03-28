package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsTime(t *testing.T) {
	a := Article{
		PublishedAt: 1,
		CreatedAt:   1,
	}
	assert.Equal(t, int64(1), a.CreatedAtAsTime().Unix())
	assert.Equal(t, int64(1), a.PublishedAtAsTime().Unix())
}

package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateAsTime(t *testing.T) {
	d := CMMeta{Date: ""}
	assert.Equal(t, int64(0), d.DateAsTime().Unix())
	d = CMMeta{Date: "2022-01-01"}
	assert.Equal(t, int64(1640995200), d.DateAsTime().Unix())
}

package bgin

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

func TestCtxUsecase(t *testing.T) {
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)

	assert.Nil(t, getCtxUsecase(ctx))
	input := usecase.Impl{}
	setCtxUsecase(ctx, &input)
	output := getCtxUsecase(ctx)
	assert.Equal(t, output, &input)
}

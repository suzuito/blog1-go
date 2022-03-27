package bgin

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/blog1-go/pkg/entity"
)

func TestCtxArticle(t *testing.T) {
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)

	assert.Nil(t, getCtxArticle(ctx))
	input := entity.Article{}
	setCtxArticle(ctx, &input)
	output := getCtxArticle(ctx)
	assert.Equal(t, output, &input)
}

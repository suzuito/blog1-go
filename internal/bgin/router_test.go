package bgin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

type testCase struct {
	Desc               string
	Input              func() *http.Request
	Setup              func(*usecase.MockUsecase)
	ExpectedStatusCode int
}

func runTest(t *testing.T, tC *testCase) {
	ctrlUsecase := gomock.NewController(t)
	mockUsecase := usecase.NewMockUsecase(ctrlUsecase)
	defer ctrlUsecase.Finish()

	if tC.Setup != nil {
		tC.Setup(mockUsecase)
	}

	r := gin.New()
	SetUpRoot(r, mockUsecase)
	w := httptest.NewRecorder()
	req := tC.Input()
	r.ServeHTTP(w, req)
	assert.Equal(t, tC.ExpectedStatusCode, w.Code)
}

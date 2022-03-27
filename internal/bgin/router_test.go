package bgin

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/blog1-go/pkg/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

type testCase struct {
	Desc               string
	Input              func() *http.Request
	Setup              func(*usecase.MockUsecase)
	ExpectedStatusCode int
}

func runTest(t *testing.T, tC *testCase) {
	dd, _ := os.Getwd()
	beforeDirPathTemplate := setting.E.DirPathTemplate
	beforeDirPathCSS := setting.E.DirPathCSS
	beforeDirPathAsset := setting.E.DirPathAsset
	gin.SetMode(gin.ReleaseMode)
	defer func() {
		setting.E.DirPathTemplate = beforeDirPathTemplate
		setting.E.DirPathCSS = beforeDirPathCSS
		setting.E.DirPathAsset = beforeDirPathAsset
		gin.SetMode(gin.DebugMode)
	}()
	setting.E.DirPathTemplate = filepath.Join(dd, "../../data/template")
	setting.E.DirPathCSS = filepath.Join(dd, "../../data/css")
	setting.E.DirPathAsset = filepath.Join(dd, "../../data/asset")

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

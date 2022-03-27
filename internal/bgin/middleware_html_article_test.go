package bgin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

func TestHTMLMiddlewareGetArticle(t *testing.T) {
	testCases := []testCase{
		{
			Input: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles/a1", nil)
				return r
			},
			Setup: func(mu *usecase.MockUsecase) {
				mu.EXPECT().GetArticle(gomock.Any(), entity.ArticleID("a1"), gomock.Any()).
					Return(errors.Wrapf(usecase.ErrNotFound, "dummy error"))
			},
			ExpectedStatusCode: http.StatusNotFound,
		},
		{
			Input: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles/a1", nil)
				return r
			},
			Setup: func(mu *usecase.MockUsecase) {
				mu.EXPECT().GetArticle(gomock.Any(), entity.ArticleID("a1"), gomock.Any()).
					Return(fmt.Errorf("dummy error"))
			},
			ExpectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.Desc, func(t *testing.T) {
			runTest(t, &tC)
		})
	}
}

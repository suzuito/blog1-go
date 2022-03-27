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

func TestHandlerHTMLGetArticle(t *testing.T) {
	testCases := []testCase{
		{
			Desc: "Success",
			Input: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles/a1", nil)
				return r
			},
			Setup: func(mu *usecase.MockUsecase) {
				mu.EXPECT().
					GetArticle(gomock.Any(), entity.ArticleID("a1"), gomock.Any()).
					SetArg(2, entity.Article{ID: "a1"})
				mu.EXPECT().
					GetArticleHTML(gomock.Any(), entity.ArticleID("a1"), gomock.Any())
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Desc: "Failed GetArticleHTML 404",
			Input: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles/a1", nil)
				return r
			},
			Setup: func(mu *usecase.MockUsecase) {
				mu.EXPECT().
					GetArticle(gomock.Any(), entity.ArticleID("a1"), gomock.Any()).
					SetArg(2, entity.Article{ID: "a1"})
				mu.EXPECT().
					GetArticleHTML(gomock.Any(), entity.ArticleID("a1"), gomock.Any()).
					Return(errors.Wrapf(usecase.ErrNotFound, "dummy error"))
			},
			ExpectedStatusCode: http.StatusNotFound,
		},
		{
			Desc: "Failed GetArticleHTML 500",
			Input: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles/a1", nil)
				return r
			},
			Setup: func(mu *usecase.MockUsecase) {
				mu.EXPECT().
					GetArticle(gomock.Any(), entity.ArticleID("a1"), gomock.Any()).
					SetArg(2, entity.Article{ID: "a1"})
				mu.EXPECT().
					GetArticleHTML(gomock.Any(), entity.ArticleID("a1"), gomock.Any()).
					Return(fmt.Errorf("dummy error"))
			},
			ExpectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tC := range testCases {
		runTest(t, &tC)
	}
}

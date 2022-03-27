package bgin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

func TestHandlerHTMLGetArticles(t *testing.T) {
	testCases := []testCase{
		{
			Desc: "Success",
			Input: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles?cursor_published_at=1&n=20", nil)
				return r
			},
			Setup: func(mu *usecase.MockUsecase) {
				mu.EXPECT().
					GetArticles(
						gomock.Any(),
						int64(1),
						"",
						usecase.CursorOrder("desc"),
						20,
						gomock.Any(),
					).
					SetArg(5, []entity.Article{
						{ID: entity.ArticleID("a1")},
					})
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Desc: "Failed GetArticles",
			Input: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles?cursor_published_at=1&n=20", nil)
				return r
			},
			Setup: func(mu *usecase.MockUsecase) {
				mu.EXPECT().
					GetArticles(
						gomock.Any(),
						int64(1),
						"",
						usecase.CursorOrder("desc"),
						20,
						gomock.Any(),
					).
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

package bgin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

func TestHandlerHTMLGetTop(t *testing.T) {
	testCases := []testCase{
		{
			Input: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			Setup: func(mu *usecase.MockUsecase) {
				mu.EXPECT().
					GetArticles(gomock.Any(), gomock.Any(), "", usecase.CursorOrderDesc, 3, gomock.Any()).
					SetArg(5, []entity.Article{})
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Input: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			Setup: func(mu *usecase.MockUsecase) {
				mu.EXPECT().
					GetArticles(gomock.Any(), gomock.Any(), "", usecase.CursorOrderDesc, 3, gomock.Any()).
					Return(fmt.Errorf("dummy errore"))
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

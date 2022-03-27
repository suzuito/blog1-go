package bgin

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

func Test(t *testing.T) {
	testCases := []testCase{
		{
			Input: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/sitemap.xml", nil)
				return r
			},
			Setup: func(mu *usecase.MockUsecase) {
				mu.EXPECT().GenerateBlogSiteMap(gomock.Any(), gomock.Any())
			},
			ExpectedStatusCode: http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.Desc, func(t *testing.T) {
			runTest(t, &tC)
		})
	}
}

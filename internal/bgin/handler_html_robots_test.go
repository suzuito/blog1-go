package bgin

import (
	"net/http"
	"testing"
)

func TestHandlerHTMLGetRobots(t *testing.T) {
	testCases := []testCase{
		{
			Input: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/robots.txt", nil)
				return r
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

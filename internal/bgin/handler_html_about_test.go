package bgin

import (
	"net/http"
	"testing"
)

func TestHandlerHTMLGetAbout(t *testing.T) {
	testCases := []testCase{
		{
			Desc: "Success",
			Input: func() *http.Request {
				r, _ := http.NewRequest(
					http.MethodGet,
					"/about",
					nil,
				)
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

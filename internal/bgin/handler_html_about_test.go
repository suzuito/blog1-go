package bgin

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/suzuito/blog1-go/pkg/setting"
)

func TestHandlerHTMLGetAbout(t *testing.T) {
	dd, _ := os.Getwd()
	setting.E.DirPathTemplate = filepath.Join(dd, "../../data/template")
	setting.E.DirPathCSS = filepath.Join(dd, "../../data/css")
	setting.E.DirPathAsset = filepath.Join(dd, "../../data/asset")
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

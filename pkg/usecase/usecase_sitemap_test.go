package usecase

import (
	"testing"
)

func TestGenerateBlogSiteMap(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// mocks, impl, closeFunc := NewMockDepends(t)
			// defer closeFunc()
		})
	}
}

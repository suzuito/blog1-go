package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/blog1-go/pkg/entity"
)

func TestGenerateBlogSiteMap(t *testing.T) {
	dummyOrigin := "https://hoge.com"
	testCases := []struct {
		desc            string
		setup           func(mocks *Mocks)
		expectedXMLURLs []XMLURL
		expectedErr     string
	}{
		{
			desc: "Success",
			setup: func(mocks *Mocks) {
				gomock.InOrder(
					mocks.DB.EXPECT().
						GetArticles(
							gomock.Any(),
							int64(0),
							"",
							CursorOrderAsc,
							100,
							gomock.Any(),
						).
						SetArg(5, []entity.Article{
							{ID: "a1"},
						}),
					mocks.DB.EXPECT().
						GetArticles(
							gomock.Any(),
							int64(0),
							"",
							CursorOrderAsc,
							100,
							gomock.Any(),
						).
						SetArg(5, []entity.Article{}),
				)
			},
			expectedXMLURLs: []XMLURL{
				{
					Loc:     dummyOrigin + "/articles/a1",
					Lastmod: "1970-01-01",
				},
				{
					Loc:     dummyOrigin + "/",
					Lastmod: "2022-03-20",
				},
				{
					Loc:     dummyOrigin + "/articles/",
					Lastmod: "2022-03-20",
				},
				{
					Loc:     dummyOrigin + "/about/",
					Lastmod: "2022-03-20",
				},
			},
		},
		{
			desc: "Cannot get articles",
			setup: func(mocks *Mocks) {
				mocks.DB.EXPECT().
					GetArticles(
						gomock.Any(),
						int64(0),
						"",
						CursorOrderAsc,
						100,
						gomock.Any(),
					).Return(fmt.Errorf("dummy error"))
			},
			expectedErr: "cannot get articles: dummy error",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			mocks, impl, closeFunc := NewMockDepends(t)
			defer closeFunc()
			tC.setup(mocks)
			real, realErr := impl.GenerateBlogSiteMap(ctx, dummyOrigin)
			if realErr != nil {
				assert.Equal(t, tC.expectedErr, realErr.Error())
				return
			}
			for _, a := range tC.expectedXMLURLs {
				assert.Contains(t, real.URLs, a)
			}
		})
	}
}

package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/blog1-go/pkg/entity"
)

func TestGetArticles(t *testing.T) {
	testCases := []struct {
		desc                   string
		setup                  func(*Mocks)
		inputCursorPublishedAt int64
		inputCursorTitle       string
		inputOrder             CursorOrder
		inputN                 int
		expectedErr            string
	}{
		{
			desc:                   "Success",
			inputCursorPublishedAt: int64(1),
			inputCursorTitle:       "hoge",
			inputOrder:             CursorOrderAsc,
			inputN:                 2,
			setup: func(mocks *Mocks) {
				mocks.DB.EXPECT().
					GetArticles(gomock.Any(), int64(1), "hoge", CursorOrder("asc"), 2, gomock.Any())
			},
		},
		{
			desc:                   "Failed",
			inputCursorPublishedAt: int64(1),
			inputCursorTitle:       "hoge",
			inputOrder:             CursorOrderAsc,
			inputN:                 2,
			expectedErr:            "dummy error",
			setup: func(mocks *Mocks) {
				mocks.DB.EXPECT().
					GetArticles(gomock.Any(), int64(1), "hoge", CursorOrder("asc"), 2, gomock.Any()).
					Return(fmt.Errorf("dummy error"))
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			mocks, impl, closeFunc := NewMockDepends(t)
			defer closeFunc()
			tC.setup(mocks)
			realArticles := []entity.Article{}
			realErr := impl.GetArticles(ctx, tC.inputCursorPublishedAt, tC.inputCursorTitle, tC.inputOrder, tC.inputN, &realArticles)
			if realErr != nil {
				assert.Equal(t, tC.expectedErr, realErr.Error())
			}
		})
	}
}

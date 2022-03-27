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

func TestGetArticle(t *testing.T) {
	testCases := []struct {
		desc           string
		setup          func(*Mocks)
		inputArticleID entity.ArticleID
		expectedErr    string
	}{
		{
			desc:           "Success",
			inputArticleID: entity.ArticleID("a01"),
			setup: func(mocks *Mocks) {
				mocks.DB.EXPECT().
					GetArticle(gomock.Any(), entity.ArticleID("a01"), gomock.Any())
			},
		},
		{
			desc:           "Failed",
			inputArticleID: entity.ArticleID("a01"),
			expectedErr:    "dummy error",
			setup: func(mocks *Mocks) {
				mocks.DB.EXPECT().
					GetArticle(gomock.Any(), entity.ArticleID("a01"), gomock.Any()).
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
			real := entity.Article{}
			err := impl.GetArticle(ctx, tC.inputArticleID, &real)
			if tC.expectedErr != "" {
				assert.NotNil(t, err)
			}
			if err != nil {
				assert.Equal(t, tC.expectedErr, err.Error())
			}
		})
	}
}

func TestGetGetArticleMarkdown(t *testing.T) {
	testCases := []struct {
		desc           string
		setup          func(*Mocks)
		inputBucket    string
		inputArticleID entity.ArticleID
		expectedErr    string
	}{
		{
			desc:           "Success",
			inputBucket:    "b1",
			inputArticleID: entity.ArticleID("a01"),
			setup: func(mocks *Mocks) {
				mocks.Storage.EXPECT().
					GetFileAsHTTPResponse(gomock.Any(), "b1/a01.md", gomock.Any(), gomock.Any())
			},
		},
		{
			desc:           "Failed",
			inputBucket:    "b1",
			inputArticleID: entity.ArticleID("a01"),
			expectedErr:    "cannot get file from b1/a01.md: dummy error",
			setup: func(mocks *Mocks) {
				mocks.Storage.EXPECT().
					GetFileAsHTTPResponse(gomock.Any(), "b1/a01.md", gomock.Any(), gomock.Any()).
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
			real := []byte{}
			err := impl.GetArticleMarkdown(ctx, tC.inputBucket, tC.inputArticleID, &real)
			if tC.expectedErr != "" {
				assert.NotNil(t, err)
			}
			if err != nil {
				assert.Equal(t, tC.expectedErr, err.Error())
			}
		})
	}
}

func TestConvertFromMarkdownToHTML(t *testing.T) {
	testCases := []struct {
		desc            string
		setup           func(*Mocks)
		inputSrcMD      []byte
		expectedHTML    string
		expectedArticle entity.Article
		expectedErr     string
	}{

		{
			desc:       "Success",
			inputSrcMD: []byte("input"),
			expectedArticle: entity.Article{
				ID: "a1",
				Images: []entity.ArticleImage{
					{URL: "img1"},
				},
				TOC: []entity.ArticleIndex{
					{ID: "idx1"},
				},
				Tags: []entity.Tag{},
			},
			expectedHTML: "modifiedHTML2",
			setup: func(mocks *Mocks) {
				mocks.MDConverter.EXPECT().
					Convert(gomock.All(), "input", gomock.Any(), gomock.Any()).
					SetArg(2, "modifiedHTML1").
					SetArg(3, CMMeta{ID: "a1"})
				mocks.HTMLEditor.EXPECT().
					ModifyHTML(gomock.Any(), "modifiedHTML1", gomock.Any()).
					SetArg(2, "modifiedHTML2")
				mocks.HTMLMediaFetcher.EXPECT().
					Fetch(gomock.Any(), "modifiedHTML2", gomock.Any()).
					SetArg(2, []entity.ArticleImage{{URL: "img1"}})
				mocks.HTMLTOCExtractor.EXPECT().
					Extract("modifiedHTML2", gomock.Any()).
					SetArg(1, []entity.ArticleIndex{{ID: "idx1"}})
			},
		},
		{
			desc:       "Failed empty ID is detected",
			inputSrcMD: []byte("input"),
			expectedArticle: entity.Article{
				Images: []entity.ArticleImage{
					{URL: "img1"},
				},
				TOC: []entity.ArticleIndex{
					{ID: "idx1"},
				},
				Tags: []entity.Tag{},
			},
			expectedHTML: "modifiedHTML2",
			setup: func(mocks *Mocks) {
				mocks.MDConverter.EXPECT().
					Convert(gomock.All(), "input", gomock.Any(), gomock.Any()).
					SetArg(2, "modifiedHTML1").
					SetArg(3, CMMeta{ID: ""})
				mocks.HTMLEditor.EXPECT().
					ModifyHTML(gomock.Any(), "modifiedHTML1", gomock.Any()).
					SetArg(2, "modifiedHTML2")
				mocks.HTMLMediaFetcher.EXPECT().
					Fetch(gomock.Any(), "modifiedHTML2", gomock.Any()).
					SetArg(2, []entity.ArticleImage{{URL: "img1"}})
				mocks.HTMLTOCExtractor.EXPECT().
					Extract("modifiedHTML2", gomock.Any()).
					SetArg(1, []entity.ArticleIndex{{ID: "idx1"}})
			},
			expectedErr: "Empty ID is detected '&{ID: Title: Description: CreatedAt:0 UpdatedAt:0 PublishedAt:0 Tags:[] Images:[{Width:0 Height:0 URL:img1 RealWidth:0 RealHeight:0}] TOC:[{ID:idx1 Name: Level:0}]}'",
		},
		{
			desc:       "Failed cannot extract toc",
			inputSrcMD: []byte("input"),
			expectedArticle: entity.Article{
				Images: []entity.ArticleImage{
					{URL: "img1"},
				},
				TOC: []entity.ArticleIndex{
					{ID: "idx1"},
				},
				Tags: []entity.Tag{},
			},
			expectedHTML: "modifiedHTML2",
			setup: func(mocks *Mocks) {
				mocks.MDConverter.EXPECT().
					Convert(gomock.All(), "input", gomock.Any(), gomock.Any()).
					SetArg(2, "modifiedHTML1").
					SetArg(3, CMMeta{ID: ""})
				mocks.HTMLEditor.EXPECT().
					ModifyHTML(gomock.Any(), "modifiedHTML1", gomock.Any()).
					SetArg(2, "modifiedHTML2")
				mocks.HTMLMediaFetcher.EXPECT().
					Fetch(gomock.Any(), "modifiedHTML2", gomock.Any()).
					SetArg(2, []entity.ArticleImage{{URL: "img1"}})
				mocks.HTMLTOCExtractor.EXPECT().
					Extract("modifiedHTML2", gomock.Any()).
					Return(fmt.Errorf("dummy error"))
			},
			expectedErr: "cannot extract toc: dummy error",
		},
		{
			desc:       "Failed cannot fetch image",
			inputSrcMD: []byte("input"),
			expectedArticle: entity.Article{
				Images: []entity.ArticleImage{
					{URL: "img1"},
				},
				TOC: []entity.ArticleIndex{
					{ID: "idx1"},
				},
				Tags: []entity.Tag{},
			},
			expectedHTML: "modifiedHTML2",
			setup: func(mocks *Mocks) {
				mocks.MDConverter.EXPECT().
					Convert(gomock.All(), "input", gomock.Any(), gomock.Any()).
					SetArg(2, "modifiedHTML1").
					SetArg(3, CMMeta{ID: ""})
				mocks.HTMLEditor.EXPECT().
					ModifyHTML(gomock.Any(), "modifiedHTML1", gomock.Any()).
					SetArg(2, "modifiedHTML2")
				mocks.HTMLMediaFetcher.EXPECT().
					Fetch(gomock.Any(), "modifiedHTML2", gomock.Any()).
					Return(fmt.Errorf("dummy error"))
			},
			expectedErr: "cannot fetch image: dummy error",
		},
		{
			desc:       "Failed cannot modify html",
			inputSrcMD: []byte("input"),
			expectedArticle: entity.Article{
				Images: []entity.ArticleImage{
					{URL: "img1"},
				},
				TOC: []entity.ArticleIndex{
					{ID: "idx1"},
				},
				Tags: []entity.Tag{},
			},
			expectedHTML: "modifiedHTML2",
			setup: func(mocks *Mocks) {
				mocks.MDConverter.EXPECT().
					Convert(gomock.All(), "input", gomock.Any(), gomock.Any()).
					SetArg(2, "modifiedHTML1").
					SetArg(3, CMMeta{ID: ""})
				mocks.HTMLEditor.EXPECT().
					ModifyHTML(gomock.Any(), "modifiedHTML1", gomock.Any()).
					Return(fmt.Errorf("dummy error"))
			},
			expectedErr: "cannot ModifyHTML: dummy error",
		},
		{
			desc:       "Failed cannot convert",
			inputSrcMD: []byte("input"),
			expectedArticle: entity.Article{
				Images: []entity.ArticleImage{
					{URL: "img1"},
				},
				TOC: []entity.ArticleIndex{
					{ID: "idx1"},
				},
				Tags: []entity.Tag{},
			},
			expectedHTML: "modifiedHTML2",
			setup: func(mocks *Mocks) {
				mocks.MDConverter.EXPECT().
					Convert(gomock.All(), "input", gomock.Any(), gomock.Any()).
					Return(fmt.Errorf("dummy error"))
			},
			expectedErr: "cannot convert: dummy error",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			mocks, impl, closeFunc := NewMockDepends(t)
			defer closeFunc()
			tC.setup(mocks)
			realHTML := ""
			realArticle := entity.Article{}
			err := impl.ConvertFromMarkdownToHTML(ctx, tC.inputSrcMD, &realHTML, &realArticle)
			if tC.expectedErr != "" {
				assert.NotNil(t, err)
			}
			if err != nil {
				assert.Equal(t, tC.expectedErr, err.Error())
				return
			}
			assert.Equal(t, tC.expectedArticle, realArticle)
			assert.Equal(t, tC.expectedHTML, realHTML)
		})
	}
}

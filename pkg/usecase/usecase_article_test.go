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
		inputArticleID entity.ArticleID
		expectedErr    string
	}{
		{
			desc:           "Success",
			inputArticleID: entity.ArticleID("a01"),
			setup: func(mocks *Mocks) {
				mocks.Storage.EXPECT().
					GetFileAsHTTPResponse(gomock.Any(), "a01.md", gomock.Any(), gomock.Any())
			},
		},
		{
			desc:           "Failed",
			inputArticleID: entity.ArticleID("a01"),
			expectedErr:    "cannot get file from a01.md: dummy error",
			setup: func(mocks *Mocks) {
				mocks.Storage.EXPECT().
					GetFileAsHTTPResponse(gomock.Any(), "a01.md", gomock.Any(), gomock.Any()).
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
			err := impl.GetArticleMarkdown(ctx, tC.inputArticleID, &real)
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

func TestUpdateArticle(t *testing.T) {
	testCases := []struct {
		desc         string
		setup        func(*Mocks)
		inputArticle entity.Article
		inputHTML    string
		expectedErr  string
	}{
		{
			desc:         "Success",
			inputArticle: entity.Article{ID: "a1"},
			inputHTML:    "input html",
			setup: func(mocks *Mocks) {
				mocks.Storage.EXPECT().
					UploadArticle(gomock.All(), gomock.Any(), "input html")
				mocks.DB.EXPECT().
					SetArticle(gomock.Any(), gomock.Any())
			},
		},
		{
			desc:         "Failed cannot set article",
			inputArticle: entity.Article{ID: "a1"},
			inputHTML:    "input html",
			setup: func(mocks *Mocks) {
				mocks.Storage.EXPECT().
					UploadArticle(gomock.All(), gomock.Any(), "input html")
				mocks.DB.EXPECT().
					SetArticle(gomock.Any(), gomock.Any()).
					Return(fmt.Errorf("dummy error"))
			},
			expectedErr: "cannot set article '&{ID:a1 Title: Description: CreatedAt:0 UpdatedAt:0 PublishedAt:0 Tags:[] Images:[] TOC:[]}': dummy error",
		},
		{
			desc:         "Failed cannot upload article",
			inputArticle: entity.Article{ID: "a1"},
			inputHTML:    "input html",
			setup: func(mocks *Mocks) {
				mocks.Storage.EXPECT().
					UploadArticle(gomock.All(), gomock.Any(), "input html").
					Return(fmt.Errorf("dummy error"))
			},
			expectedErr: "cannot upload article '&{ID:a1 Title: Description: CreatedAt:0 UpdatedAt:0 PublishedAt:0 Tags:[] Images:[] TOC:[]}': dummy error",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			mocks, impl, closeFunc := NewMockDepends(t)
			defer closeFunc()
			tC.setup(mocks)
			err := impl.UpdateArticle(ctx, &tC.inputArticle, tC.inputHTML)
			if tC.expectedErr != "" {
				assert.NotNil(t, err)
			}
			if err != nil {
				assert.Equal(t, tC.expectedErr, err.Error())
				return
			}
		})
	}
}

func TestDeleteArticle(t *testing.T) {
	testCases := []struct {
		desc           string
		setup          func(*Mocks)
		inputArticleID entity.ArticleID
		expectedErr    string
	}{
		{
			desc:           "Success",
			inputArticleID: "a1",
			setup: func(mocks *Mocks) {
				mocks.Storage.EXPECT().
					DeleteArticle(gomock.Any(), entity.ArticleID("a1"))
				mocks.DB.EXPECT().
					DeleteArticle(gomock.Any(), entity.ArticleID("a1"))
			},
		},
		{
			desc:           "Failed cannot delete article",
			inputArticleID: "a1",
			setup: func(mocks *Mocks) {
				mocks.Storage.EXPECT().
					DeleteArticle(gomock.Any(), entity.ArticleID("a1"))
				mocks.DB.EXPECT().
					DeleteArticle(gomock.Any(), entity.ArticleID("a1")).
					Return(fmt.Errorf("dummy error"))
			},
			expectedErr: "cannot delete article 'a1': dummy error",
		},
		{
			desc:           "Failed cannot delete article",
			inputArticleID: "a1",
			setup: func(mocks *Mocks) {
				mocks.Storage.EXPECT().
					DeleteArticle(gomock.All(), entity.ArticleID("a1")).
					Return(fmt.Errorf("dummy error"))
			},
			expectedErr: "cannot delete article 'a1': dummy error",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			mocks, impl, closeFunc := NewMockDepends(t)
			defer closeFunc()
			tC.setup(mocks)
			err := impl.DeleteArticle(ctx, tC.inputArticleID)
			if tC.expectedErr != "" {
				assert.NotNil(t, err)
			}
			if err != nil {
				assert.Equal(t, tC.expectedErr, err.Error())
				return
			}
		})
	}
}

func TestGetArticleHTML(t *testing.T) {
	testCases := []struct {
		desc           string
		setup          func(*Mocks)
		inputArticleID entity.ArticleID
		expectedErr    string
	}{
		{
			desc:           "Success",
			inputArticleID: "a1",
			setup: func(mocks *Mocks) {
				mocks.Storage.EXPECT().
					GetFileAsHTTPResponse(gomock.Any(), "a1.html", gomock.Any(), gomock.Any())
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			mocks, impl, closeFunc := NewMockDepends(t)
			defer closeFunc()
			tC.setup(mocks)
			err := impl.GetArticleHTML(ctx, tC.inputArticleID, &[]byte{})
			if tC.expectedErr != "" {
				assert.NotNil(t, err)
			}
			if err != nil {
				assert.Equal(t, tC.expectedErr, err.Error())
				return
			}
		})
	}
}

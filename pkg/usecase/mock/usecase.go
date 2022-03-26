// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/suzuito/blog1-go/pkg/entity"
	usecase "github.com/suzuito/blog1-go/pkg/usecase"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// ConvertFromMarkdownToHTML mocks base method.
func (m *MockUsecase) ConvertFromMarkdownToHTML(ctx context.Context, srcMD []byte, retHTML *string, article *entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertFromMarkdownToHTML", ctx, srcMD, retHTML, article)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConvertFromMarkdownToHTML indicates an expected call of ConvertFromMarkdownToHTML.
func (mr *MockUsecaseMockRecorder) ConvertFromMarkdownToHTML(ctx, srcMD, retHTML, article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertFromMarkdownToHTML", reflect.TypeOf((*MockUsecase)(nil).ConvertFromMarkdownToHTML), ctx, srcMD, retHTML, article)
}

// DeleteArticle mocks base method.
func (m *MockUsecase) DeleteArticle(ctx context.Context, articleID entity.ArticleID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArticle", ctx, articleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArticle indicates an expected call of DeleteArticle.
func (mr *MockUsecaseMockRecorder) DeleteArticle(ctx, articleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArticle", reflect.TypeOf((*MockUsecase)(nil).DeleteArticle), ctx, articleID)
}

// GenerateBlogSiteMap mocks base method.
func (m *MockUsecase) GenerateBlogSiteMap(ctx context.Context, origin string) (*usecase.XMLURLSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateBlogSiteMap", ctx, origin)
	ret0, _ := ret[0].(*usecase.XMLURLSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateBlogSiteMap indicates an expected call of GenerateBlogSiteMap.
func (mr *MockUsecaseMockRecorder) GenerateBlogSiteMap(ctx, origin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateBlogSiteMap", reflect.TypeOf((*MockUsecase)(nil).GenerateBlogSiteMap), ctx, origin)
}

// GetArticle mocks base method.
func (m *MockUsecase) GetArticle(ctx context.Context, id entity.ArticleID, article *entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticle", ctx, id, article)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticle indicates an expected call of GetArticle.
func (mr *MockUsecaseMockRecorder) GetArticle(ctx, id, article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticle", reflect.TypeOf((*MockUsecase)(nil).GetArticle), ctx, id, article)
}

// GetArticleHTML mocks base method.
func (m *MockUsecase) GetArticleHTML(ctx context.Context, id entity.ArticleID, body *[]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleHTML", ctx, id, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticleHTML indicates an expected call of GetArticleHTML.
func (mr *MockUsecaseMockRecorder) GetArticleHTML(ctx, id, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleHTML", reflect.TypeOf((*MockUsecase)(nil).GetArticleHTML), ctx, id, body)
}

// GetArticleMarkdown mocks base method.
func (m *MockUsecase) GetArticleMarkdown(ctx context.Context, bucket string, articleID entity.ArticleID, dst *[]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleMarkdown", ctx, bucket, articleID, dst)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticleMarkdown indicates an expected call of GetArticleMarkdown.
func (mr *MockUsecaseMockRecorder) GetArticleMarkdown(ctx, bucket, articleID, dst interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleMarkdown", reflect.TypeOf((*MockUsecase)(nil).GetArticleMarkdown), ctx, bucket, articleID, dst)
}

// GetArticles mocks base method.
func (m *MockUsecase) GetArticles(ctx context.Context, cursorPublishedAt int64, cursorTitle string, order usecase.CursorOrder, n int, articles *[]entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles", ctx, cursorPublishedAt, cursorTitle, order, n, articles)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockUsecaseMockRecorder) GetArticles(ctx, cursorPublishedAt, cursorTitle, order, n, articles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockUsecase)(nil).GetArticles), ctx, cursorPublishedAt, cursorTitle, order, n, articles)
}

// UpdateArticle mocks base method.
func (m *MockUsecase) UpdateArticle(ctx context.Context, article *entity.Article, htmlString string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArticle", ctx, article, htmlString)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateArticle indicates an expected call of UpdateArticle.
func (mr *MockUsecaseMockRecorder) UpdateArticle(ctx, article, htmlString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArticle", reflect.TypeOf((*MockUsecase)(nil).UpdateArticle), ctx, article, htmlString)
}

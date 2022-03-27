// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/db.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/suzuito/blog1-go/pkg/entity"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// DeleteArticle mocks base method.
func (m *MockDB) DeleteArticle(ctx context.Context, articleID entity.ArticleID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArticle", ctx, articleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArticle indicates an expected call of DeleteArticle.
func (mr *MockDBMockRecorder) DeleteArticle(ctx, articleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArticle", reflect.TypeOf((*MockDB)(nil).DeleteArticle), ctx, articleID)
}

// GetArticle mocks base method.
func (m *MockDB) GetArticle(ctx context.Context, articleID entity.ArticleID, article *entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticle", ctx, articleID, article)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticle indicates an expected call of GetArticle.
func (mr *MockDBMockRecorder) GetArticle(ctx, articleID, article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticle", reflect.TypeOf((*MockDB)(nil).GetArticle), ctx, articleID, article)
}

// GetArticles mocks base method.
func (m *MockDB) GetArticles(ctx context.Context, cursorPublishedAt int64, cursorTitle string, order CursorOrder, n int, articles *[]entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles", ctx, cursorPublishedAt, cursorTitle, order, n, articles)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockDBMockRecorder) GetArticles(ctx, cursorPublishedAt, cursorTitle, order, n, articles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockDB)(nil).GetArticles), ctx, cursorPublishedAt, cursorTitle, order, n, articles)
}

// SetArticle mocks base method.
func (m *MockDB) SetArticle(ctx context.Context, article *entity.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetArticle", ctx, article)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetArticle indicates an expected call of SetArticle.
func (mr *MockDBMockRecorder) SetArticle(ctx, article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetArticle", reflect.TypeOf((*MockDB)(nil).SetArticle), ctx, article)
}

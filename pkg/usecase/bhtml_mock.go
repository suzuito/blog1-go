// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/bhtml.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/suzuito/blog1-go/pkg/entity"
)

// MockHTMLEditor is a mock of HTMLEditor interface.
type MockHTMLEditor struct {
	ctrl     *gomock.Controller
	recorder *MockHTMLEditorMockRecorder
}

// MockHTMLEditorMockRecorder is the mock recorder for MockHTMLEditor.
type MockHTMLEditorMockRecorder struct {
	mock *MockHTMLEditor
}

// NewMockHTMLEditor creates a new mock instance.
func NewMockHTMLEditor(ctrl *gomock.Controller) *MockHTMLEditor {
	mock := &MockHTMLEditor{ctrl: ctrl}
	mock.recorder = &MockHTMLEditorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTMLEditor) EXPECT() *MockHTMLEditorMockRecorder {
	return m.recorder
}

// ModifyHTML mocks base method.
func (m *MockHTMLEditor) ModifyHTML(ctx context.Context, src string, dst *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyHTML", ctx, src, dst)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyHTML indicates an expected call of ModifyHTML.
func (mr *MockHTMLEditorMockRecorder) ModifyHTML(ctx, src, dst interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyHTML", reflect.TypeOf((*MockHTMLEditor)(nil).ModifyHTML), ctx, src, dst)
}

// MockHTMLMediaFetcher is a mock of HTMLMediaFetcher interface.
type MockHTMLMediaFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockHTMLMediaFetcherMockRecorder
}

// MockHTMLMediaFetcherMockRecorder is the mock recorder for MockHTMLMediaFetcher.
type MockHTMLMediaFetcherMockRecorder struct {
	mock *MockHTMLMediaFetcher
}

// NewMockHTMLMediaFetcher creates a new mock instance.
func NewMockHTMLMediaFetcher(ctrl *gomock.Controller) *MockHTMLMediaFetcher {
	mock := &MockHTMLMediaFetcher{ctrl: ctrl}
	mock.recorder = &MockHTMLMediaFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTMLMediaFetcher) EXPECT() *MockHTMLMediaFetcherMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockHTMLMediaFetcher) Fetch(ctx context.Context, src string, images *[]entity.ArticleImage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, src, images)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fetch indicates an expected call of Fetch.
func (mr *MockHTMLMediaFetcherMockRecorder) Fetch(ctx, src, images interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockHTMLMediaFetcher)(nil).Fetch), ctx, src, images)
}

// MockHTMLTOCExtractor is a mock of HTMLTOCExtractor interface.
type MockHTMLTOCExtractor struct {
	ctrl     *gomock.Controller
	recorder *MockHTMLTOCExtractorMockRecorder
}

// MockHTMLTOCExtractorMockRecorder is the mock recorder for MockHTMLTOCExtractor.
type MockHTMLTOCExtractorMockRecorder struct {
	mock *MockHTMLTOCExtractor
}

// NewMockHTMLTOCExtractor creates a new mock instance.
func NewMockHTMLTOCExtractor(ctrl *gomock.Controller) *MockHTMLTOCExtractor {
	mock := &MockHTMLTOCExtractor{ctrl: ctrl}
	mock.recorder = &MockHTMLTOCExtractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTMLTOCExtractor) EXPECT() *MockHTMLTOCExtractorMockRecorder {
	return m.recorder
}

// Extract mocks base method.
func (m *MockHTMLTOCExtractor) Extract(src string, idx *[]entity.ArticleIndex) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Extract", src, idx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Extract indicates an expected call of Extract.
func (mr *MockHTMLTOCExtractorMockRecorder) Extract(src, idx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Extract", reflect.TypeOf((*MockHTMLTOCExtractor)(nil).Extract), src, idx)
}

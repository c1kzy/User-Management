// Code generated by MockGen. DO NOT EDIT.
// Source: restapi/internal/infrastructure/inputports/grpc/handler (interfaces: ImplementedPostService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	post "restapi/internal/infrastructure/inputports/grpc/proto/post"

	gomock "github.com/golang/mock/gomock"
)

// ImplementedPostService is a mock of ImplementedPostService interface.
type ImplementedPostService struct {
	ctrl     *gomock.Controller
	recorder *ImplementedPostServiceMockRecorder
}

// ImplementedPostServiceMockRecorder is the mock recorder for ImplementedPostService.
type ImplementedPostServiceMockRecorder struct {
	mock *ImplementedPostService
}

// NewImplementedPostService creates a new mock instance.
func NewImplementedPostService(ctrl *gomock.Controller) *ImplementedPostService {
	mock := &ImplementedPostService{ctrl: ctrl}
	mock.recorder = &ImplementedPostServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *ImplementedPostService) EXPECT() *ImplementedPostServiceMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *ImplementedPostService) CreatePost(arg0 context.Context, arg1 *post.PostRequest) (*post.PostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(*post.PostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *ImplementedPostServiceMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*ImplementedPostService)(nil).CreatePost), arg0, arg1)
}

// DeletePost mocks base method.
func (m *ImplementedPostService) DeletePost(arg0 context.Context, arg1 *post.PostRequest) (*post.PostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0, arg1)
	ret0, _ := ret[0].(*post.PostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePost indicates an expected call of DeletePost.
func (mr *ImplementedPostServiceMockRecorder) DeletePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*ImplementedPostService)(nil).DeletePost), arg0, arg1)
}

// ListPosts mocks base method.
func (m *ImplementedPostService) ListPosts(arg0 context.Context, arg1 *post.Page) (*post.PublicPostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPosts", arg0, arg1)
	ret0, _ := ret[0].(*post.PublicPostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPosts indicates an expected call of ListPosts.
func (mr *ImplementedPostServiceMockRecorder) ListPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPosts", reflect.TypeOf((*ImplementedPostService)(nil).ListPosts), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *ImplementedPostService) UpdatePost(arg0 context.Context, arg1 *post.PostRequest) (*post.PostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(*post.PostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *ImplementedPostServiceMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*ImplementedPostService)(nil).UpdatePost), arg0, arg1)
}

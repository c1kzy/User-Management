// Code generated by MockGen. DO NOT EDIT.
// Source: restapi/internal/domain (interfaces: PostService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	domain "restapi/internal/domain"

	gomock "github.com/golang/mock/gomock"
)

// PostService is a mock of PostService interface.
type PostService struct {
	ctrl     *gomock.Controller
	recorder *PostServiceMockRecorder
}

// PostServiceMockRecorder is the mock recorder for PostService.
type PostServiceMockRecorder struct {
	mock *PostService
}

// NewPostService creates a new mock instance.
func NewPostService(ctrl *gomock.Controller) *PostService {
	mock := &PostService{ctrl: ctrl}
	mock.recorder = &PostServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *PostService) EXPECT() *PostServiceMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *PostService) CreatePost(arg0 domain.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePost indicates an expected call of CreatePost.
func (mr *PostServiceMockRecorder) CreatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*PostService)(nil).CreatePost), arg0)
}

// DeletePost mocks base method.
func (m *PostService) DeletePost(arg0 int, arg1 domain.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *PostServiceMockRecorder) DeletePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*PostService)(nil).DeletePost), arg0, arg1)
}

// GetPost mocks base method.
func (m *PostService) GetPost(arg0 int) (domain.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0)
	ret0, _ := ret[0].(domain.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *PostServiceMockRecorder) GetPost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*PostService)(nil).GetPost), arg0)
}

// GetUserID mocks base method.
func (m *PostService) GetUserID(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserID", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserID indicates an expected call of GetUserID.
func (mr *PostServiceMockRecorder) GetUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserID", reflect.TypeOf((*PostService)(nil).GetUserID), arg0)
}

// ListPosts mocks base method.
func (m *PostService) ListPosts(arg0 string) ([]domain.PublicPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPosts", arg0)
	ret0, _ := ret[0].([]domain.PublicPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPosts indicates an expected call of ListPosts.
func (mr *PostServiceMockRecorder) ListPosts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPosts", reflect.TypeOf((*PostService)(nil).ListPosts), arg0)
}

// UpdatePost mocks base method.
func (m *PostService) UpdatePost(arg0 int, arg1 domain.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *PostServiceMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*PostService)(nil).UpdatePost), arg0, arg1)
}

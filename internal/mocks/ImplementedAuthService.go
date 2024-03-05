// Code generated by MockGen. DO NOT EDIT.
// Source: restapi/internal/infrastructure/inputports/grpc/handler (interfaces: ImplementedAuthService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	auth "restapi/internal/infrastructure/inputports/grpc/proto/auth"

	gomock "github.com/golang/mock/gomock"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// ImplementedAuthService is a mock of ImplementedAuthService interface.
type ImplementedAuthService struct {
	ctrl     *gomock.Controller
	recorder *ImplementedAuthServiceMockRecorder
}

// ImplementedAuthServiceMockRecorder is the mock recorder for ImplementedAuthService.
type ImplementedAuthServiceMockRecorder struct {
	mock *ImplementedAuthService
}

// NewImplementedAuthService creates a new mock instance.
func NewImplementedAuthService(ctrl *gomock.Controller) *ImplementedAuthService {
	mock := &ImplementedAuthService{ctrl: ctrl}
	mock.recorder = &ImplementedAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *ImplementedAuthService) EXPECT() *ImplementedAuthServiceMockRecorder {
	return m.recorder
}

// SignIn mocks base method.
func (m *ImplementedAuthService) SignIn(arg0 context.Context, arg1 *auth.SignInRequest) (*auth.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", arg0, arg1)
	ret0, _ := ret[0].(*auth.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *ImplementedAuthServiceMockRecorder) SignIn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*ImplementedAuthService)(nil).SignIn), arg0, arg1)
}

// SignUp mocks base method.
func (m *ImplementedAuthService) SignUp(arg0 context.Context, arg1 *auth.User) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *ImplementedAuthServiceMockRecorder) SignUp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*ImplementedAuthService)(nil).SignUp), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: restapi/internal/domain (interfaces: RatingService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	domain "restapi/internal/domain"

	gomock "github.com/golang/mock/gomock"
)

// RatingService is a mock of RatingService interface.
type RatingService struct {
	ctrl     *gomock.Controller
	recorder *RatingServiceMockRecorder
}

// RatingServiceMockRecorder is the mock recorder for RatingService.
type RatingServiceMockRecorder struct {
	mock *RatingService
}

// NewRatingService creates a new mock instance.
func NewRatingService(ctrl *gomock.Controller) *RatingService {
	mock := &RatingService{ctrl: ctrl}
	mock.recorder = &RatingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *RatingService) EXPECT() *RatingServiceMockRecorder {
	return m.recorder
}

// GetVotes mocks base method.
func (m *RatingService) GetVotes(arg0, arg1 int) (domain.Ratings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVotes", arg0, arg1)
	ret0, _ := ret[0].(domain.Ratings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVotes indicates an expected call of GetVotes.
func (mr *RatingServiceMockRecorder) GetVotes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotes", reflect.TypeOf((*RatingService)(nil).GetVotes), arg0, arg1)
}

// InsertVote mocks base method.
func (m *RatingService) InsertVote(arg0, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertVote", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertVote indicates an expected call of InsertVote.
func (mr *RatingServiceMockRecorder) InsertVote(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertVote", reflect.TypeOf((*RatingService)(nil).InsertVote), arg0, arg1, arg2)
}

// UpdateVote mocks base method.
func (m *RatingService) UpdateVote(arg0 domain.UserRating, arg1 domain.Ratings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVote", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVote indicates an expected call of UpdateVote.
func (mr *RatingServiceMockRecorder) UpdateVote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVote", reflect.TypeOf((*RatingService)(nil).UpdateVote), arg0, arg1)
}

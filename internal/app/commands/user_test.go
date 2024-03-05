package commands

import (
	"fmt"
	"restapi/internal/domain"
	"restapi/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserDB_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	userService := mocks.NewUserService(ctrl)

	tests := []struct {
		name          string
		mocks         func(service *mocks.UserService)
		user          domain.User
		expectedError error
	}{
		{
			name: "created",
			user: domain.User{
				FirstName: "test",
				LastName:  "testenko",
				Email:     "test@test.com",
				Password:  "test",
			},
			mocks: func(service *mocks.UserService) {
				service.EXPECT().Create(domain.User{
					FirstName: "test",
					LastName:  "testenko",
					Email:     "test@test.com",
					Password:  "test",
				}).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "empty",
			user: domain.User{},
			mocks: func(service *mocks.UserService) {
				service.EXPECT().Create(domain.User{}).Return(fmt.Errorf("user fields are empty"))
			},
			expectedError: fmt.Errorf("user fields are empty"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(userService)
			err := userService.Create(tc.user)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

func TestUserDB_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	userService := mocks.NewUserService(ctrl)

	tests := []struct {
		name          string
		mocks         func(service *mocks.UserService)
		id            int
		expectedError error
	}{
		{
			name: "deleted",
			mocks: func(service *mocks.UserService) {
				service.EXPECT().Delete(1).Return(nil)
			},
			id:            1,
			expectedError: nil,
		},
		{
			name: "invalid id",
			mocks: func(service *mocks.UserService) {
				service.EXPECT().Delete(0).Return(fmt.Errorf("invalid id"))
			},
			expectedError: fmt.Errorf("invalid id"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(userService)
			err := userService.Delete(tc.id)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

func TestUserDB_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	userService := mocks.NewUserService(ctrl)

	tests := []struct {
		name          string
		mocks         func(service *mocks.UserService)
		id            int
		user          domain.User
		want          domain.User
		expectedError error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.UserService) {
				service.EXPECT().Update(1, domain.User{
					FirstName: "test",
					LastName:  "test",
					Email:     "test@test.com",
					Password:  "test",
				}).Return(nil)
			},
			id: 1,
			user: domain.User{
				FirstName: "test",
				LastName:  "test",
				Email:     "test@test.com",
				Password:  "test",
			},
			expectedError: nil,
		},
		{
			name: "not ok",
			mocks: func(service *mocks.UserService) {
				service.EXPECT().Update(0, domain.User{
					FirstName: "test",
					LastName:  "test",
					Email:     "test@test.com",
					Password:  "test",
				}).Return(fmt.Errorf("invalid id"))
			},
			id: 0,
			user: domain.User{
				FirstName: "test",
				LastName:  "test",
				Email:     "test@test.com",
				Password:  "test",
			},
			expectedError: fmt.Errorf("invalid id"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(userService)
			err := userService.Update(tc.id, tc.user)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

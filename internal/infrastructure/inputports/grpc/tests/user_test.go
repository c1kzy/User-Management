package handler

import (
	"context"
	"fmt"
	"restapi/internal/infrastructure/inputports/grpc/proto/user"
	"restapi/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserAPI_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userService := mocks.NewImplementedUserService(ctrl)

	ctx := context.Background()

	tests := []struct {
		name        string
		mocks       func(service *mocks.ImplementedUserService)
		req         *user.UpdateUserRequest
		want        *user.UserResponse
		expectedErr error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.ImplementedUserService) {
				service.EXPECT().UpdateUser(ctx, &user.UpdateUserRequest{
					ID:        1,
					FirstName: "test",
					LastName:  "test",
					Email:     "test@test.com",
					Password:  "test",
					Role:      1,
				}).Return(&user.UserResponse{Message: "user updated"}, nil)
			},
			req: &user.UpdateUserRequest{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				Email:     "test@test.com",
				Password:  "test",
				Role:      1,
			},
			want:        &user.UserResponse{Message: "user updated"},
			expectedErr: nil,
		},
		{
			name: "invalid id",
			mocks: func(service *mocks.ImplementedUserService) {
				service.EXPECT().UpdateUser(ctx, &user.UpdateUserRequest{
					ID:        0,
					FirstName: "test",
					LastName:  "test",
					Email:     "test@test.com",
					Password:  "test",
					Role:      1,
				}).Return(nil, fmt.Errorf("invalid id"))
			},
			req: &user.UpdateUserRequest{
				ID:        0,
				FirstName: "test",
				LastName:  "test",
				Email:     "test@test.com",
				Password:  "test",
				Role:      1,
			},
			want:        nil,
			expectedErr: fmt.Errorf("invalid id"),
		},
		{
			name: "empty first name",
			mocks: func(service *mocks.ImplementedUserService) {
				service.EXPECT().UpdateUser(ctx, &user.UpdateUserRequest{
					ID:        1,
					FirstName: "",
					LastName:  "test",
					Email:     "1234@test.com",
					Password:  "test",
					Role:      1,
				}).Return(nil, fmt.Errorf("first name is required"))
			},
			req: &user.UpdateUserRequest{
				ID:        1,
				FirstName: "",
				LastName:  "test",
				Email:     "1234@test.com",
				Password:  "test",
				Role:      1,
			},
			want:        nil,
			expectedErr: fmt.Errorf("first name is required"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(userService)
			got, err := userService.UpdateUser(ctx, tt.req)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserAPI_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userService := mocks.NewImplementedUserService(ctrl)

	ctx := context.Background()

	tests := []struct {
		name        string
		mocks       func(service *mocks.ImplementedUserService)
		req         *user.DeleteUserRequest
		want        *user.UserResponse
		expectedErr error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.ImplementedUserService) {
				service.EXPECT().DeleteUser(ctx, &user.DeleteUserRequest{Id: 1}).Return(
					&user.UserResponse{Message: "user deleted"}, nil)
			},
			req:         &user.DeleteUserRequest{Id: 1},
			want:        &user.UserResponse{Message: "user deleted"},
			expectedErr: nil,
		},
		{
			name: "invalid id",
			mocks: func(service *mocks.ImplementedUserService) {
				service.EXPECT().DeleteUser(ctx, &user.DeleteUserRequest{Id: 0}).Return(
					nil, fmt.Errorf("invalid id"))
			},
			req:         &user.DeleteUserRequest{Id: 0},
			want:        nil,
			expectedErr: fmt.Errorf("invalid id"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(userService)
			got, err := userService.DeleteUser(ctx, tt.req)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

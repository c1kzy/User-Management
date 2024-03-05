package handler

import (
	"context"
	"fmt"
	"restapi/internal/infrastructure/inputports/grpc/proto/auth"
	"restapi/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestAuthAPI_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	authService := mocks.NewImplementedAuthService(ctrl)

	ctx := context.Background()

	tests := []struct {
		name        string
		mocks       func(service *mocks.ImplementedAuthService)
		req         *auth.User
		want        *emptypb.Empty
		expectedErr error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.ImplementedAuthService) {
				service.EXPECT().SignUp(ctx, &auth.User{
					FirstName: "test",
					LastName:  "test2",
					Email:     "test@test.com",
					Password:  "test",
					Role:      1,
				}).Return(&emptypb.Empty{}, nil)
			},
			req: &auth.User{
				FirstName: "test",
				LastName:  "test2",
				Email:     "test@test.com",
				Password:  "test",
				Role:      1,
			},
			want:        nil,
			expectedErr: nil,
		},
		{
			name: "empty last name",
			mocks: func(service *mocks.ImplementedAuthService) {
				service.EXPECT().SignUp(ctx, &auth.User{
					FirstName: "test",
					LastName:  "",
					Email:     "test@test.com",
					Password:  "",
					Role:      1,
				}).Return(&emptypb.Empty{}, fmt.Errorf("last name is required"))
			},
			req: &auth.User{
				FirstName: "test",
				LastName:  "",
				Email:     "test@test.com",
				Password:  "",
				Role:      1,
			},
			want:        nil,
			expectedErr: fmt.Errorf("last name is required"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(authService)
			got, err := authService.SignUp(ctx, tt.req)
			assert.Equal(t, err, tt.expectedErr)
			assert.Equal(t, &emptypb.Empty{}, got)
		})
	}
}

func TestAuthAPI_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	authService := mocks.NewImplementedAuthService(ctrl)

	ctx := context.Background()

	tests := []struct {
		name        string
		mocks       func(service *mocks.ImplementedAuthService)
		req         *auth.SignInRequest
		want        *auth.LoginResponse
		expectedErr error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.ImplementedAuthService) {
				service.EXPECT().SignIn(ctx, &auth.SignInRequest{
					Email:    "test123@test.com",
					Password: "test",
				}).Return(&auth.LoginResponse{Token: "123"}, nil)
			},
			req: &auth.SignInRequest{
				Email:    "test123@test.com",
				Password: "test",
			},
			want:        &auth.LoginResponse{Token: "123"},
			expectedErr: nil,
		},
		{
			name: "empty email",
			mocks: func(service *mocks.ImplementedAuthService) {
				service.EXPECT().SignIn(ctx, &auth.SignInRequest{
					Email:    "",
					Password: "test",
				}).Return(nil, fmt.Errorf("email is required"))
			},
			req: &auth.SignInRequest{
				Email:    "",
				Password: "test",
			},
			want:        nil,
			expectedErr: fmt.Errorf("email is required"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(authService)
			got, err := authService.SignIn(ctx, tt.req)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

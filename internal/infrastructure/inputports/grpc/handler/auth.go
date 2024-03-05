package handler

import (
	"context"
	"fmt"
	"restapi/internal"
	"restapi/internal/app/commands"
	"restapi/internal/domain"
	"restapi/internal/infrastructure/inputports/grpc/proto/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ImplementedAuthService interface {
	SignUp(ctx context.Context, req *auth.User) (*emptypb.Empty, error)
	SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.LoginResponse, error)
}
type AuthAPI struct {
	auth.UnimplementedAuthServiceServer
	Service     *internal.Service
	VoteService *commands.VoteService
}

func RegisterAuth(s *grpc.Server, service *internal.Service, voteService *commands.VoteService) {
	auth.RegisterAuthServiceServer(s, &AuthAPI{Service: service, VoteService: voteService})
}

func (a *AuthAPI) SignUp(ctx context.Context, req *auth.User) (*emptypb.Empty, error) {
	newUser := domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      int(req.Role),
	}

	if err := newUser.IsValid(); err != nil {
		return nil, err
	}

	err := a.Service.UserService.Create(newUser)
	if err != nil {
		return nil, ResponseWithError(codes.Internal, err)
	}

	return nil, nil

}

func (a *AuthAPI) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.LoginResponse, error) {
	if req.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	if req.Password == "" {
		return nil, fmt.Errorf("password is required")
	}

	token, err := a.Service.UserService.GenerateToken(req.Email, req.Password)
	if err != nil {
		return nil, ResponseWithError(codes.Internal, err)
	}

	return &auth.LoginResponse{Token: token}, nil
}

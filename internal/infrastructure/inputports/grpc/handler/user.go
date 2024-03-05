package handler

import (
	"context"
	"restapi/internal"
	"restapi/internal/app/commands"
	"restapi/internal/domain"
	"restapi/internal/infrastructure/inputports/grpc/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type ImplementedUserService interface {
	UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UserResponse, error)
	DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.UserResponse, error)
}
type UserAPI struct {
	user.UnimplementedUserServiceServer
	service     *internal.Service
	voteService *commands.VoteService
}

func RegisterUser(s *grpc.Server, service *internal.Service, voteService *commands.VoteService) {
	user.RegisterUserServiceServer(s, &UserAPI{service: service, voteService: voteService})
}

func (u *UserAPI) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UserResponse, error) {
	id := int(req.ID)

	if id == 0 {
		return nil, invalidIDErr
	}

	userToUpdate := domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      int(req.Role),
	}

	if err := userToUpdate.IsValid(); err != nil {
		return nil, err
	}

	err := u.service.UserService.Update(id, userToUpdate)
	if err != nil {
		return nil, ResponseWithError(codes.Internal, err)
	}

	return &user.UserResponse{Message: userUpdated}, nil
}

func (u *UserAPI) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.UserResponse, error) {
	id := int(req.Id)
	if id == 0 {
		return nil, invalidIDErr
	}

	err := u.service.UserService.Delete(id)
	if err != nil {
		return nil, ResponseWithError(codes.Internal, err)
	}

	return &user.UserResponse{Message: userDeleted}, nil
}

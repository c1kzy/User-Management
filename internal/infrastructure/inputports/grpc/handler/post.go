package handler

import (
	"context"
	"fmt"
	"restapi/internal"
	"restapi/internal/app/commands"
	"restapi/internal/domain"
	"restapi/internal/infrastructure/inputports/grpc/proto/post"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type ImplementedPostService interface {
	CreatePost(ctx context.Context, req *post.PostRequest) (*post.PostResponse, error)
	UpdatePost(ctx context.Context, req *post.PostRequest) (*post.PostResponse, error)
	DeletePost(ctx context.Context, req *post.PostRequest) (*post.PostResponse, error)
	ListPosts(ctx context.Context, req *post.Page) (*post.PublicPostResponse, error)
}
type PostAPI struct {
	post.UnimplementedPostServiceServer
	service     *internal.Service
	voteService *commands.VoteService
}

func RegisterPost(s *grpc.Server, service *internal.Service, voteService *commands.VoteService) {
	post.RegisterPostServiceServer(s, &PostAPI{service: service, voteService: voteService})
}

func (p *PostAPI) CreatePost(ctx context.Context, req *post.PostRequest) (*post.PostResponse, error) {
	newPost := domain.Post{
		UserID: int(req.UserId),
		Name:   req.Name,
		Text:   req.Text,
	}

	if newPost.UserID == 0 {
		return nil, fmt.Errorf("userID is required")
	}

	if err := newPost.IsValid(); err != nil {
		return nil, err
	}

	err := p.service.PostService.CreatePost(newPost)
	if err != nil {
		return nil, ResponseWithError(codes.Internal, err)
	}

	return &post.PostResponse{Message: postCreated}, nil
}

func (p *PostAPI) UpdatePost(ctx context.Context, req *post.PostRequest) (*post.PostResponse, error) {
	id := int(req.ID)

	if id == 0 {
		return nil, invalidIDErr
	}

	postToUpdate := domain.Post{
		Name: req.Name,
		Text: req.Text,
	}

	if err := postToUpdate.IsValid(); err != nil {
		return nil, err
	}

	err := p.service.PostService.UpdatePost(id, postToUpdate)
	if err != nil {
		return nil, ResponseWithError(codes.Internal, err)
	}

	return &post.PostResponse{Message: postUpdated}, nil
}

func (p *PostAPI) DeletePost(ctx context.Context, req *post.PostRequest) (*post.PostResponse, error) {
	id := int(req.ID)

	postToDelete := domain.Post{
		UserID: int(req.UserId),
	}

	if id == 0 || req.UserId == 0 {
		return nil, fmt.Errorf("invalid post or user ID")
	}

	err := p.service.PostService.DeletePost(id, postToDelete)
	if err != nil {
		return nil, ResponseWithError(codes.Internal, err)
	}

	return &post.PostResponse{Message: postDeleted}, nil
}

func (p *PostAPI) ListPosts(ctx context.Context, req *post.Page) (*post.PublicPostResponse, error) {
	page := req.Page

	if page == "" {
		return nil, invalidPageErr
	}

	posts, err := p.service.PostService.ListPosts(page)
	if err != nil {
		return nil, ResponseWithError(codes.Internal, err)
	}

	result := make([]*post.PublicPost, len(posts))
	for i := range posts {
		result[i] = &post.PublicPost{
			ID:   int32(posts[i].ID),
			Name: posts[i].Name,
			Text: posts[i].Text,
		}
	}

	return &post.PublicPostResponse{PublicPosts: result}, nil
}

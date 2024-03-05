package handler

import (
	"context"
	"fmt"
	"restapi/internal/infrastructure/inputports/grpc/proto/post"
	"restapi/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPostAPI_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	postService := mocks.NewImplementedPostService(ctrl)

	ctx := context.Background()
	tests := []struct {
		name        string
		mocks       func(service *mocks.ImplementedPostService)
		req         *post.PostRequest
		want        *post.PostResponse
		expectedErr error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.ImplementedPostService) {
				service.EXPECT().CreatePost(ctx, &post.PostRequest{
					UserId: 1,
					Name:   "test",
					Text:   "test",
				}).Return(&post.PostResponse{Message: "post created"}, nil)
			},
			req: &post.PostRequest{
				UserId: 1,
				Name:   "test",
				Text:   "test",
			},
			want:        &post.PostResponse{Message: "post created"},
			expectedErr: nil,
		},
		{
			name: "empty post name",
			mocks: func(service *mocks.ImplementedPostService) {
				service.EXPECT().CreatePost(ctx, &post.PostRequest{
					UserId: 1,
					Name:   "",
					Text:   "test",
				}).Return(nil, fmt.Errorf("post name is required"))
			},
			req: &post.PostRequest{
				UserId: 1,
				Name:   "",
				Text:   "test",
			},
			want:        nil,
			expectedErr: fmt.Errorf("post name is required"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(postService)
			got, err := postService.CreatePost(ctx, tt.req)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPostAPI_UpdatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	postService := mocks.NewImplementedPostService(ctrl)

	ctx := context.Background()
	tests := []struct {
		name        string
		mocks       func(service *mocks.ImplementedPostService)
		req         *post.PostRequest
		want        *post.PostResponse
		expectedErr error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.ImplementedPostService) {
				service.EXPECT().UpdatePost(ctx, &post.PostRequest{
					ID:   1,
					Name: "test",
					Text: "test",
				}).Return(&post.PostResponse{Message: "post updated"}, nil)
			},
			req: &post.PostRequest{
				ID:   1,
				Name: "test",
				Text: "test",
			},
			want:        &post.PostResponse{Message: "post updated"},
			expectedErr: nil,
		},
		{
			name: "invalid id",
			mocks: func(service *mocks.ImplementedPostService) {
				service.EXPECT().UpdatePost(ctx, &post.PostRequest{
					ID:   0,
					Name: "test",
					Text: "test",
				}).Return(nil, fmt.Errorf("invalid id"))
			},
			req: &post.PostRequest{
				ID:   0,
				Name: "test",
				Text: "test",
			},
			want:        nil,
			expectedErr: fmt.Errorf("invalid id"),
		},
		{
			name: "empty name",
			mocks: func(service *mocks.ImplementedPostService) {
				service.EXPECT().UpdatePost(ctx, &post.PostRequest{
					ID:   1,
					Name: "",
					Text: "test",
				}).Return(nil, fmt.Errorf("post name is required"))
			},
			req: &post.PostRequest{
				ID:   1,
				Name: "",
				Text: "test",
			},
			want:        nil,
			expectedErr: fmt.Errorf("post name is required"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(postService)
			got, err := postService.UpdatePost(ctx, tt.req)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPostAPI_DeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	postService := mocks.NewImplementedPostService(ctrl)

	ctx := context.Background()
	tests := []struct {
		name        string
		mocks       func(service *mocks.ImplementedPostService)
		req         *post.PostRequest
		want        *post.PostResponse
		expectedErr error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.ImplementedPostService) {
				service.EXPECT().DeletePost(ctx, &post.PostRequest{
					ID:     1,
					UserId: 1,
				}).Return(&post.PostResponse{Message: "post deleted"}, nil)
			},
			req: &post.PostRequest{
				ID:     1,
				UserId: 1,
			},
			want:        &post.PostResponse{Message: "post deleted"},
			expectedErr: nil,
		},
		{
			name: "invalid post id",
			mocks: func(service *mocks.ImplementedPostService) {
				service.EXPECT().DeletePost(ctx, &post.PostRequest{
					ID:     0,
					UserId: 1,
				}).Return(nil, fmt.Errorf("invalid post or user ID"))
			},
			req: &post.PostRequest{
				ID:     0,
				UserId: 1,
			},
			want:        nil,
			expectedErr: fmt.Errorf("invalid post or user ID"),
		},
		{
			name: "invalid user id",
			mocks: func(service *mocks.ImplementedPostService) {
				service.EXPECT().DeletePost(ctx, &post.PostRequest{
					ID:     1,
					UserId: 0,
				}).Return(nil, fmt.Errorf("fields are empty"))
			},
			req: &post.PostRequest{
				ID:     1,
				UserId: 0,
			},
			want:        nil,
			expectedErr: fmt.Errorf("fields are empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(postService)
			got, err := postService.DeletePost(ctx, tt.req)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPostAPI_ListPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	postService := mocks.NewImplementedPostService(ctrl)

	ctx := context.Background()
	tests := []struct {
		name        string
		mocks       func(service *mocks.ImplementedPostService)
		req         *post.Page
		want        *post.PublicPostResponse
		expectedErr error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.ImplementedPostService) {
				service.EXPECT().ListPosts(ctx, &post.Page{Page: "1"}).Return(&post.PublicPostResponse{PublicPosts: []*post.PublicPost{
					{
						ID:   1,
						Name: "test",
						Text: "test",
					},
					{
						ID:   2,
						Name: "test2",
						Text: "test2",
					},
				}}, nil)
			},
			req: &post.Page{Page: "1"},
			want: &post.PublicPostResponse{PublicPosts: []*post.PublicPost{
				{
					ID:   1,
					Name: "test",
					Text: "test",
				},
				{
					ID:   2,
					Name: "test2",
					Text: "test2",
				},
			}},
			expectedErr: nil,
		},
		{
			name: "invalid page",
			mocks: func(service *mocks.ImplementedPostService) {
				service.EXPECT().ListPosts(ctx, &post.Page{Page: ""}).Return(nil, fmt.Errorf("invalid page"))
			},
			req:         &post.Page{Page: ""},
			want:        nil,
			expectedErr: fmt.Errorf("invalid page"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(postService)
			got, err := postService.ListPosts(ctx, tt.req)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

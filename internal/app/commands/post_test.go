package commands

import (
	"fmt"
	"restapi/internal/domain"
	"restapi/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPostDB_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	postService := mocks.NewPostService(ctrl)

	tests := []struct {
		name          string
		mocks         func(service *mocks.PostService)
		post          domain.Post
		expectedError error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.PostService) {
				service.EXPECT().CreatePost(domain.Post{
					UserID: 1,
					Text:   "test",
				}).Return(nil)
			},
			post: domain.Post{
				UserID: 1,
				Text:   "test",
			},
			expectedError: nil,
		},
		{
			name: "empty",
			mocks: func(service *mocks.PostService) {
				service.EXPECT().CreatePost(domain.Post{}).Return(fmt.Errorf("post fields are empty"))
			},
			post:          domain.Post{},
			expectedError: fmt.Errorf("post fields are empty"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(postService)
			err := postService.CreatePost(tc.post)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

func TestPostDB_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	postService := mocks.NewPostService(ctrl)

	tests := []struct {
		name          string
		mocks         func(service *mocks.PostService)
		post          domain.Post
		id            int
		expectedError error
	}{
		{
			name: "deleted",
			mocks: func(service *mocks.PostService) {
				service.EXPECT().DeletePost(1, domain.Post{
					UserID: 1,
				}).Return(nil)
			},
			post: domain.Post{
				UserID: 1,
			},
			id:            1,
			expectedError: nil,
		},
		{
			name: "invalid id",
			mocks: func(service *mocks.PostService) {
				service.EXPECT().DeletePost(0, domain.Post{}).Return(fmt.Errorf("invalid id"))
			},
			expectedError: fmt.Errorf("invalid id"),
		},
		{
			name: "invalid user id",
			mocks: func(service *mocks.PostService) {
				service.EXPECT().DeletePost(1, domain.Post{
					UserID: 0,
				}).Return(fmt.Errorf("invalid id"))
			},
			id:            1,
			expectedError: fmt.Errorf("invalid id"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(postService)
			err := postService.DeletePost(tc.id, tc.post)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

func TestPostDB_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	postService := mocks.NewPostService(ctrl)

	tests := []struct {
		name          string
		mocks         func(service *mocks.PostService)
		id            int
		post          domain.Post
		want          domain.Post
		expectedError error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.PostService) {
				service.EXPECT().UpdatePost(1, domain.Post{
					Text: "test update",
				}).Return(nil)
			},
			id: 1,
			post: domain.Post{
				Text: "test update",
			},
			expectedError: nil,
		},
		{
			name: "not ok",
			mocks: func(service *mocks.PostService) {
				service.EXPECT().UpdatePost(0, domain.Post{
					Text: "test update",
				}).Return(fmt.Errorf("invalid id"))
			},
			id: 0,
			post: domain.Post{
				Text: "test update",
			},
			expectedError: fmt.Errorf("invalid id"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(postService)
			err := postService.UpdatePost(tc.id, tc.post)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

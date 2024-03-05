package handler

import (
	"context"
	"fmt"
	"restapi/internal/infrastructure/inputports/grpc/proto/rating"
	"restapi/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRatingAPI_UserVote(t *testing.T) {
	ctrl := gomock.NewController(t)
	ratingService := mocks.NewImplementedRatingService(ctrl)

	ctx := context.Background()
	tests := []struct {
		name        string
		mocks       func(service *mocks.ImplementedRatingService)
		req         *rating.UserRatingRequest
		want        *rating.RatingResponse
		expectedErr error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.ImplementedRatingService) {
				service.EXPECT().UserVote(ctx, &rating.UserRatingRequest{
					Id:     1,
					UserId: 1,
					PostId: 1,
					Vote:   1,
				}).Return(&rating.RatingResponse{Message: "user voted"}, nil)
			},
			req: &rating.UserRatingRequest{
				Id:     1,
				UserId: 1,
				PostId: 1,
				Vote:   1,
			},
			want:        &rating.RatingResponse{Message: "user voted"},
			expectedErr: nil,
		},
		{
			name: "invalid userID",
			mocks: func(service *mocks.ImplementedRatingService) {
				service.EXPECT().UserVote(ctx, &rating.UserRatingRequest{
					Id:     1,
					UserId: 0,
					PostId: 1,
					Vote:   1,
				}).Return(nil, fmt.Errorf("invalid userID or postID"))
			},
			req: &rating.UserRatingRequest{
				Id:     1,
				UserId: 0,
				PostId: 1,
				Vote:   1,
			},
			want:        nil,
			expectedErr: fmt.Errorf("invalid userID or postID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(ratingService)
			got, err := ratingService.UserVote(ctx, tt.req)
			assert.Equal(t, err, tt.expectedErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

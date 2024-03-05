package commands

import (
	"fmt"
	"restapi/internal/domain"
	"restapi/internal/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRatingRepository_GetVotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	ratingService := mocks.NewRatingService(ctrl)

	tests := []struct {
		name           string
		mocks          func(service *mocks.RatingService)
		userID         int
		postID         int
		expectedRating domain.Ratings
		expectedError  error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.RatingService) {
				service.EXPECT().GetVotes(1, 1).Return(domain.Ratings{
					ID:         1,
					FromUserID: 1,
					ToPostID:   1,
					UserVote:   1,
					WhenVoted:  time.Time{},
					Status:     1,
				}, nil)
			},
			userID: 1,
			postID: 1,
			expectedRating: domain.Ratings{
				ID:         1,
				FromUserID: 1,
				ToPostID:   1,
				UserVote:   1,
				WhenVoted:  time.Time{},
				Status:     1,
			},
			expectedError: nil,
		},
		{
			name: "not ok",
			mocks: func(service *mocks.RatingService) {
				service.EXPECT().GetVotes(1, 1).Return(domain.Ratings{}, fmt.Errorf("unable to get votes from DB. See error:"))
			},
			userID:         1,
			postID:         1,
			expectedRating: domain.Ratings{},
			expectedError:  fmt.Errorf("unable to get votes from DB. See error:"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(ratingService)
			rating, err := ratingService.GetVotes(tc.userID, tc.postID)
			assert.Equal(t, rating, tc.expectedRating)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

func TestRatingRepository_InsertVote(t *testing.T) {
	ctrl := gomock.NewController(t)
	ratingService := mocks.NewRatingService(ctrl)

	tests := []struct {
		name          string
		mocks         func(service *mocks.RatingService)
		userID        int
		postID        int
		userVote      int
		expectedError error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.RatingService) {
				service.EXPECT().InsertVote(1, 1, 1).Return(nil)
			},
			userID:        1,
			postID:        1,
			userVote:      1,
			expectedError: nil,
		},
		{
			name: "not ok",
			mocks: func(service *mocks.RatingService) {
				service.EXPECT().InsertVote(3, 2, 1).Return(fmt.Errorf("vote error occured. See error:"))
			},
			userID:        3,
			postID:        2,
			userVote:      1,
			expectedError: fmt.Errorf("vote error occured. See error:"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(ratingService)
			err := ratingService.InsertVote(tc.userID, tc.postID, tc.userVote)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

func TestRatingRepository_UpdateVote(t *testing.T) {
	ctrl := gomock.NewController(t)
	ratingService := mocks.NewRatingService(ctrl)

	tests := []struct {
		name          string
		mocks         func(service *mocks.RatingService)
		input         domain.UserRating
		rating        domain.Ratings
		expectedError error
	}{
		{
			name: "ok",
			mocks: func(service *mocks.RatingService) {
				service.EXPECT().UpdateVote(domain.UserRating{
					ID:   1,
					Vote: -1,
				}, domain.Ratings{
					UserVote: 1,
					Status:   1,
				}).Return(nil)
			},
			input: domain.UserRating{
				ID:   1,
				Vote: -1,
			},
			rating: domain.Ratings{
				UserVote: 1,
				Status:   1,
			},
			expectedError: nil,
		},
		{
			name: "empty fields",
			mocks: func(service *mocks.RatingService) {
				service.EXPECT().UpdateVote(domain.UserRating{}, domain.Ratings{
					UserVote: 1,
					Status:   1,
				}).Return(fmt.Errorf("ratings fields are empty"))
			},
			input: domain.UserRating{},
			rating: domain.Ratings{
				UserVote: 1,
				Status:   1,
			},
			expectedError: fmt.Errorf("ratings fields are empty"),
		},
		{
			name: "empty ID",
			mocks: func(service *mocks.RatingService) {
				service.EXPECT().UpdateVote(domain.UserRating{
					Vote: -1,
				}, domain.Ratings{
					UserVote: 1,
					Status:   1,
				}).Return(fmt.Errorf("please specify ID to update"))
			},
			input: domain.UserRating{
				Vote: -1,
			},
			rating: domain.Ratings{
				UserVote: 1,
				Status:   1,
			},
			expectedError: fmt.Errorf("please specify ID to update"),
		},
		{
			name: "vote twice",
			mocks: func(service *mocks.RatingService) {
				service.EXPECT().UpdateVote(domain.UserRating{
					ID:   1,
					Vote: 1,
				}, domain.Ratings{
					UserVote: 1,
					Status:   1,
				}).Return(fmt.Errorf("cannot vote twice"))
			},
			input: domain.UserRating{
				ID:   1,
				Vote: 1,
			},
			rating: domain.Ratings{
				UserVote: 1,
				Status:   1,
			},
			expectedError: fmt.Errorf("cannot vote twice"),
		},
		{
			name: "db error",
			mocks: func(service *mocks.RatingService) {
				service.EXPECT().UpdateVote(domain.UserRating{
					ID:   1,
					Vote: 1,
				}, domain.Ratings{
					UserVote: -1,
					Status:   1,
				}).Return(fmt.Errorf("update vote error occured. See error:"))
			},
			input: domain.UserRating{
				ID:   1,
				Vote: 1,
			},
			rating: domain.Ratings{
				UserVote: -1,
				Status:   1,
			},
			expectedError: fmt.Errorf("update vote error occured. See error:"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks(ratingService)
			err := ratingService.UpdateVote(tc.input, tc.rating)
			assert.Equal(t, err, tc.expectedError)
		})
	}
}

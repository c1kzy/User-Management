package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"restapi/internal"
	"restapi/internal/app/commands"
	"restapi/internal/domain"
	"restapi/internal/infrastructure/inputports/grpc/proto/rating"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type ImplementedRatingService interface {
	UserVote(ctx context.Context, req *rating.UserRatingRequest) (*rating.RatingResponse, error)
}
type RatingAPI struct {
	rating.UnimplementedRatingServiceServer
	service     *internal.Service
	voteService *commands.VoteService
}

func RegisterRating(s *grpc.Server, service *internal.Service, voteService *commands.VoteService) {
	rating.RegisterRatingServiceServer(s, &RatingAPI{service: service, voteService: voteService})
}

func (r *RatingAPI) UserVote(ctx context.Context, req *rating.UserRatingRequest) (*rating.RatingResponse, error) {
	currentTime := time.Now().UTC()

	newRating := domain.UserRating{
		ID:   int(req.Id),
		Vote: int(req.Vote),
	}

	userID := int(req.UserId)
	postID := int(req.PostId)

	if userID == 0 || postID == 0 {
		return nil, fmt.Errorf("invalid userID:%v or postID:%v", userID, postID)
	}

	postUserID, err := r.service.PostService.GetUserID(int(req.PostId))
	if err != nil {
		return nil, ResponseWithError(codes.Internal, err)
	}

	if postUserID == int(req.UserId) {
		return nil, ResponseWithError(codes.Internal, selfVoteErr)
	}

	votes, err := r.service.RatingService.GetVotes(userID, postID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, ResponseWithError(codes.Internal, err)
		}
	}

	postVote := domain.VoteSumUpdate{
		Vote:   newRating.Vote,
		PostID: postID,
	}

	if votes.FromUserID == userID && votes.ToPostID == postID {
		if !currentTime.After(votes.WhenVoted.Add(1 * time.Second)) {
			return nil, ResponseWithError(codes.Internal, cannotVoteErr)
		}

		if newRating.Vote == votes.UserVote {
			postVote = domain.VoteSumUpdate{
				Vote:   -newRating.Vote,
				PostID: postID,
			}
		}

		if newRating.Vote == votes.UserVote {
			postVote = domain.VoteSumUpdate{
				Vote:   -newRating.Vote,
				PostID: postID,
			}
		}

		err = r.service.RatingService.UpdateVote(newRating, votes)
		if err != nil {
			return nil, ResponseWithError(codes.Internal, err)
		}

		err = r.service.VoteService.UpdatePostVotes(postVote)
		if err != nil {
			return nil, ResponseWithError(codes.Internal, err)
		}

		return &rating.RatingResponse{Message: voteUpdated}, nil
	}

	err = r.service.RatingService.InsertVote(userID, postID, newRating.Vote)
	if err != nil {
		return nil, ResponseWithError(codes.Internal, err)

	}

	err = r.service.VoteService.UpdatePostVotes(postVote)
	if err != nil {
		return nil, ResponseWithError(codes.Internal, err)
	}

	return &rating.RatingResponse{Message: userVoted}, nil

}

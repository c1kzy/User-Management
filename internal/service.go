package internal

import (
	"restapi/internal/app/cache"
	"restapi/internal/app/commands"
	"restapi/internal/domain"
	"restapi/internal/pkg/database"
)

type Service struct {
	domain.UserService
	domain.PostService
	domain.RatingService
	domain.VoteService
}

func NewService(db *database.DB, cfg *database.Config, client *cache.Redis) *Service {
	return &Service{

		UserService:   commands.NewUserRepository(db, cfg, client),
		PostService:   commands.NewPostRepository(db, client),
		RatingService: commands.NewRatingRepository(db, client),
		VoteService:   commands.NewVoteService(db),
	}
}

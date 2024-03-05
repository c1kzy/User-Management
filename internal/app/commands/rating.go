package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"restapi/internal/app/cache"
	"restapi/internal/domain"
	"restapi/internal/pkg/database"
	"strconv"
	"time"

	"github.com/phuslu/log"
	"github.com/redis/go-redis/v9"
)

type RatingRepository struct {
	db     *database.DB
	client *cache.Redis
}

func NewRatingRepository(db *database.DB, client *cache.Redis) *RatingRepository {
	return &RatingRepository{db: db, client: client}

}

const (
	NotVoted = iota
	Voted
)

func (r *RatingRepository) InsertVote(userID, postID, userVote int) error {
	_, err := r.db.Exec("call insertVote($1, $2, $3, $4, $5)", userID, postID, userVote, time.Now().UTC(), Voted)
	if err != nil {
		return fmt.Errorf("vote error occured. See error:%w", err)
	}

	return nil
}

func (r *RatingRepository) UpdateVote(input domain.UserRating, rating domain.Ratings) error {
	status := Voted

	if input.ID == 0 {
		return fmt.Errorf("please specify ID to update")
	}

	if reflect.DeepEqual(input, domain.UserRating{}) || reflect.DeepEqual(rating, domain.Ratings{}) {
		return fmt.Errorf("ratings fields are empty")
	}

	vote := input.Vote

	if input.Vote == rating.UserVote {
		vote = input.Vote - rating.UserVote
		status = NotVoted
	}

	_, err := r.db.Exec("call updateVote($1, $2, $3, $4)", vote, time.Now().UTC(), status, input.ID)
	if err != nil {
		return fmt.Errorf("update vote error occured. See error:%w", err)
	}

	return nil
}

func (r *RatingRepository) GetVotes(userID, postID int) (domain.Ratings, error) {
	id := strconv.Itoa(postID)

	val, err := r.client.Get(ctx, id)
	if !errors.Is(err, redis.Nil) {
		log.Error().Err(err)
	}

	var ratings domain.Ratings
	if err := json.Unmarshal([]byte(val), &ratings); err != nil {
		query := "SELECT * FROM ratings WHERE from_user_id=$1 AND to_post_id=$2"

		result, err := r.db.GetVoteFromDB(query, userID, postID)
		if err != nil {
			log.Error().Err(err)
			return domain.Ratings{}, fmt.Errorf("unable to get votes from DB. See error:%w", err)
		}

		dbRating, ok := result.(domain.Ratings)
		if !ok {
			return domain.Ratings{}, fmt.Errorf("returned result is not a type of Votes struct. See result:%v", result)
		}

		r.client.Set(ctx, id, dbRating, 1*time.Minute)

		return dbRating, nil
	}

	return ratings, err

}

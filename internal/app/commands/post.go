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

	"github.com/google/uuid"
	"github.com/phuslu/log"
	"github.com/redis/go-redis/v9"
)

type PostRepository struct {
	db     *database.DB
	client *cache.Redis
}

func NewPostRepository(db *database.DB, client *cache.Redis) *PostRepository {
	return &PostRepository{db: db, client: client}
}

func (p *PostRepository) CreatePost(post domain.Post) error {
	if reflect.DeepEqual(post, domain.Post{}) {
		return fmt.Errorf("post fields are empty")
	}

	_, err := p.db.Exec("call insertPost($1, $2, $3, $4, $5, $6)", post.UserID, post.Name, post.Text, time.Now().UTC(), uuid.New(), created)

	return err
}
func (p *PostRepository) UpdatePost(id int, post domain.Post) error {
	if id == 0 {
		return fmt.Errorf("invalid id")
	}

	_, err := p.db.Exec("call updatePost($1, $2, $3, $4)", time.Now().UTC(), post.Text, updated, id)

	return err
}
func (p *PostRepository) DeletePost(id int, post domain.Post) error {
	if id == 0 && post.UserID == 0 {
		return fmt.Errorf("invalid id")
	}

	_, err := p.db.Exec("call deletePost($1, $2)", deleted, id)

	return err
}

func (p *PostRepository) GetPost(id int) (domain.Post, error) {
	if id == 0 {
		return domain.Post{}, fmt.Errorf("invalid id")
	}

	var post domain.Post
	postID := strconv.Itoa(id)

	val, err := p.client.Get(ctx, postID)
	if !errors.Is(err, redis.Nil) {
		log.Error().Err(err)
	}

	if err := json.Unmarshal([]byte(val), &post); err != nil {
		log.Error().Err(err)
		query := "SELECT * FROM posts WHERE id=$1"

		result, err := p.db.GetPostFromDB(query, id)
		if err != nil {
			log.Error().Err(err)
			return domain.Post{}, fmt.Errorf("unable to get post from post DB. See error:%w", err)
		}

		dbPost, ok := result.(domain.Post)
		if !ok {
			return domain.Post{}, fmt.Errorf("returned result is not a type of Post struct. See result:%v", result)
		}

		p.client.Set(ctx, postID, dbPost, 1*time.Minute)

		return dbPost, nil
	}

	return post, err

}

func (p *PostRepository) ListPosts(page string) ([]domain.PublicPost, error) {
	offset := 10
	pageInt, _ := strconv.Atoi(page)

	query := fmt.Sprintf("SELECT id, name, text FROM posts WHERE status <> 3 LIMIT 10 OFFSET %v", (pageInt-1)*offset)

	val, err := p.client.Get(ctx, page)
	if !errors.Is(err, redis.Nil) {
		log.Error().Err(err)
	}

	var posts []domain.PublicPost
	if err := json.Unmarshal([]byte(val), &posts); err != nil {
		result, err := p.db.GetAllPosts(query)
		if err != nil {
			log.Error().Err(err)
			return []domain.PublicPost{}, err
		}

		posts = result.([]domain.PublicPost)

		p.client.Set(ctx, page, posts, 1*time.Minute)

		return posts, nil
	}

	return posts, nil
}

func (p *PostRepository) GetUserID(postID int) (int, error) {
	query := "SELECT user_id FROM posts WHERE id =$1"

	id := strconv.Itoa(postID)

	val, err := p.client.Get(ctx, id)
	if !errors.Is(err, redis.Nil) {
		log.Error().Err(err)
	}

	var userID int
	if err := json.Unmarshal([]byte(val), &userID); err != nil {
		result, err := p.db.GetPostFromDB(query, postID)
		if err != nil {
			return 0, fmt.Errorf("unable to get userID from post DB. See error:%w", err)
		}

		post, ok := result.(domain.Post)
		if !ok {
			return 0, fmt.Errorf("returned result is not a type of Posts struct. See result:%v", result)
		}

		userID = post.UserID

		p.client.Set(ctx, id, userID, 1*time.Minute)

		return userID, nil
	}

	return userID, err
}

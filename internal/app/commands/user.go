package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"restapi/internal/app/auth"
	"restapi/internal/app/cache"
	"restapi/internal/domain"
	"restapi/internal/pkg/database"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/phuslu/log"
	"github.com/redis/go-redis/v9"
)

type UserRepository struct {
	db     *database.DB
	cfg    *database.Config
	client *cache.Redis
}

var ctx = context.Background()

func NewUserRepository(db *database.DB, cfg *database.Config, client *cache.Redis) *UserRepository {
	return &UserRepository{db: db, cfg: cfg, client: client}
}

func (u *UserRepository) Create(user domain.User) error {
	if reflect.DeepEqual(user, domain.User{}) {
		return fmt.Errorf("user fields are empty")
	}

	user.Password = auth.GeneratePasswordHash(user.Password, u.cfg.Salt)

	_, err := u.db.Exec("call insertUser($1, $2, $3, $4, $5, $6, $7, $8)", user.FirstName, user.LastName, user.Email, user.Password, user.Role, time.Now().UTC(), uuid.New(), created)

	return err
}

func (u *UserRepository) Update(id int, user domain.User) error {
	if id == 0 {
		return fmt.Errorf("invalid id")
	}

	user.Password = auth.GeneratePasswordHash(user.Password, u.cfg.Salt)

	_, err := u.db.Exec("call updateUser($1, $2, $3, $4, $5, $6, $7, $8)", user.FirstName, user.LastName, user.Email, user.Password, user.Role, time.Now().UTC(), updated, id)

	return err
}

func (u *UserRepository) Delete(id int) error {
	if id == 0 {
		return fmt.Errorf("invalid id")
	}

	_, err := u.db.Exec("call deleteUser($1,$2)", deleted, id)

	return err
}

func (u *UserRepository) GetUser(email, password string) (domain.User, error) {
	if email == "" || password == "" {
		return domain.User{}, fmt.Errorf("invalid credentials")
	}

	val, err := u.client.Get(ctx, email)
	if !errors.Is(err, redis.Nil) {
		log.Error().Err(err)
	}

	var user domain.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		query := "SELECT * FROM users WHERE email=$1 AND hashed_password=$2"

		result, err := u.db.GetUserFromDB(query, email, password)
		if err != nil {
			log.Error().Err(err)
			return domain.User{}, fmt.Errorf("unable to get user from DB. See error:%w", err)
		}

		dbUser, ok := result.(domain.User)
		if !ok {
			log.Error().Err(fmt.Errorf("returned result is not a type of User struct. See result:%v", result))
			return domain.User{}, fmt.Errorf("returned result is not a type of User struct. See result:%v", result)
		}

		u.client.Set(ctx, email, dbUser, 1*time.Minute)

		return dbUser, nil
	}

	return user, err

}

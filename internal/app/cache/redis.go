package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/phuslu/log"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	redis *redis.Client
}

type RedisService interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, v any, expiration time.Duration)
}

func NewRedisDB(cfg *RedisConfig) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Pass,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	return &Redis{redis: client}, nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, key).Result()
}

func (r *Redis) Set(ctx context.Context, key string, v any, expiration time.Duration) {
	valueMarshal, err := json.Marshal(v)
	if err != nil {
		log.Error().Err(err)
		return
	}

	err = r.redis.Set(ctx, key, valueMarshal, expiration).Err()
	if err != nil {
		log.Error().Err(err)
	}

}

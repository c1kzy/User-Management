package main

import (
	"fmt"
	"restapi/internal"

	"restapi/internal/app/cache"
	"restapi/internal/app/commands"
	"restapi/internal/infrastructure/inputports/grpc/server"

	"restapi/internal/pkg/database"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/phuslu/log"
)

// @title User Management API
// @version 1.0
// @description API Server for user and profiles management

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey apiKeyAuth
// @in header
// @name Authorization

func main() {
	cfg := &database.Config{}
	redisCfg := &cache.RedisConfig{}

	internal.NewLogger()

	if envErr := godotenv.Load(); envErr != nil {
		log.Fatal().Err(fmt.Errorf("error loading .env file: %w", envErr))
	}

	if err := env.Parse(cfg); err != nil {
		log.Error().Err(err)
		return
	}

	if err := env.Parse(redisCfg); err != nil {
		log.Error().Err(err)
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal().Err(err)
	}

	voteService := commands.NewVoteService(db)

	redisClient, err := cache.NewRedisDB(redisCfg)
	if err != nil {
		log.Fatal().Err(err)
	}

	service := internal.NewService(db, cfg, redisClient)

	srv := new(server.Server)

	if err := srv.Run(service, voteService); err != nil {
		log.Fatal().Err(err)
		return
	}

}

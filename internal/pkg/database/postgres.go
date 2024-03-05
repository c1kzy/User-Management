package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/phuslu/log"
)

type DB struct {
	Db *sqlx.DB
}

func NewPostgresDB(cfg *Config) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	return &DB{
			Db: db,
		},
		nil
}

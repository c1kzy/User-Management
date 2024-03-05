package database

import "restapi/internal/domain"

func (db *DB) Exec(query string, args ...any) (interface{}, error) {
	return db.Db.Exec(query, args...)

}

func (db *DB) GetUserFromDB(query string, args ...interface{}) (interface{}, error) {
	var dest domain.User
	err := db.Db.Get(&dest, query, args...)

	return dest, err
}

func (db *DB) GetPostFromDB(query string, args ...interface{}) (interface{}, error) {
	var dest domain.Post
	err := db.Db.Get(&dest, query, args...)

	return dest, err
}

func (db *DB) GetAllPosts(query string, args ...interface{}) (interface{}, error) {
	var dest []domain.PublicPost
	err := db.Db.Select(&dest, query, args...)

	return dest, err
}

func (db *DB) GetVoteFromDB(query string, args ...interface{}) (interface{}, error) {
	var dest domain.Ratings
	err := db.Db.Get(&dest, query, args...)

	return dest, err
}

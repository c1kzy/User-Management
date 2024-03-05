package database

type DBService interface {
	Exec(query string, args ...any) (interface{}, error)
	GetUserFromDB(query string, args ...interface{}) (interface{}, error)
	GetPostFromDB(query string, args ...interface{}) (interface{}, error)
	GetAllPosts(query string, args ...interface{}) ([]interface{}, error)
}

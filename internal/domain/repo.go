package domain

type UserService interface {
	GetUser(email, password string) (User, error)
	Create(user User) error
	Update(id int, user User) error
	Delete(id int) error
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (*TokenClaims, error)
}

type PostService interface {
	CreatePost(post Post) error
	UpdatePost(id int, post Post) error
	DeletePost(id int, post Post) error
	GetPost(id int) (Post, error)
	ListPosts(page string) ([]PublicPost, error)
	GetUserID(userID int) (int, error)
}

type RatingService interface {
	InsertVote(userID, postID, userVote int) error
	UpdateVote(input UserRating, ratings Ratings) error
	GetVotes(userID, postID int) (Ratings, error)
}

type VoteService interface {
	UpdatePostVotes(vote VoteSumUpdate) error
}

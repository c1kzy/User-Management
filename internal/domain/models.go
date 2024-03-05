package domain

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Post struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Text      string    `json:"text" db:"text"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	OID       uuid.UUID `json:"oID" db:"oid"`
	Vote      int       `json:"vote_sum" db:"vote_sum"`
	Status    int       `json:"status" db:"status"`
}

type User struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"hashed_password"`
	Role      int       `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	OID       uuid.UUID `json:"oID" db:"oid"`
	Status    int       `json:"status" db:"status"`
}

type PublicPost struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Text string `json:"text" db:"text"`
}

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID int `json:"user_id"`
	Role   int `json:"role"`
}

type Ratings struct {
	ID         int       `json:"id" db:"id"`
	FromUserID int       `json:"from_user_id" db:"from_user_id"`
	ToPostID   int       `json:"to_post_id" db:"to_post_id"`
	UserVote   int       `json:"user_vote" db:"user_vote"`
	WhenVoted  time.Time `json:"when_voted" db:"when_voted"`
	Status     int       `json:"status" db:"status"`
}

type UserRating struct {
	ID   int `json:"id" db:"id"`
	Vote int `json:"vote"`
}

type VoteSumUpdate struct {
	Vote   int
	PostID int
}

func (u *User) IsValid() error {
	if u.FirstName == "" {
		return fmt.Errorf("first name is required")
	}
	if u.LastName == "" {
		return fmt.Errorf("last name is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}
	if u.Role == 0 {
		return fmt.Errorf("role is required")
	}

	return nil
}

func (p *Post) IsValid() error {
	if p.Name == "" {
		return fmt.Errorf("post name is required")
	}
	if p.Text == "" {
		return fmt.Errorf("text is required")
	}

	return nil
}

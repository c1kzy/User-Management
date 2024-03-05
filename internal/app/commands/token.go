package commands

import (
	"fmt"
	"restapi/internal/app/auth"
	"restapi/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (u *UserRepository) GenerateToken(email, password string) (string, error) {
	user, err := u.GetUser(email, auth.GeneratePasswordHash(password, u.cfg.Salt))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &domain.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(24 * time.Hour)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		UserID: user.ID,
		Role:   user.Role,
	})

	return token.SignedString([]byte(u.cfg.SigningKey))
}

func (u *UserRepository) ParseToken(accessToken string) (*domain.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(u.cfg.SigningKey), nil
	})

	if err != nil {
		return &domain.TokenClaims{}, err
	}

	claims, ok := token.Claims.(*domain.TokenClaims)
	if !ok {
		return &domain.TokenClaims{}, fmt.Errorf("invalid type of token claims")
	}

	return claims, nil
}

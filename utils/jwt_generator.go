package utils

import (
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewJwtToken(user model.User, ttl time.Duration, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["login"] = user.Username
	claims["exp"] = time.Now().Add(ttl).Unix()
	claims["role"] = user.Role

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

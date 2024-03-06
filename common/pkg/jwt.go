package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {
	secret := []byte("secret string")

	tokenCliams := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iss":      time.Now(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := tokenCliams.SignedString(secret)
	return token, err
}

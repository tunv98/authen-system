package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userId, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	return token.SignedString(secretKey)
}

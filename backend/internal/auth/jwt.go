package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Secret = []byte("SUPER_SECRET_KEY")

func GenerateToken(userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}

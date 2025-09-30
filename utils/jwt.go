package utils

import (
	"help_desk/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user models.User) (string, time.Time, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	expiresAt := time.Now().Add(30 * time.Minute)
	claims := jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"exp":      expiresAt.Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", time.Time{}, err
	}
	return signedToken, expiresAt, nil
}

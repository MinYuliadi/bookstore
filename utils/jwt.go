package utils

import (
	"bookstore/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtKey = []byte("")

func GenerateJWT(username string) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}

	jwtKey = []byte(os.Getenv("JWT_KEY"))

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := models.JWTClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*models.JWTClaim, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	jwtKey = []byte(os.Getenv("JWT_KEY"))

	claims := &models.JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

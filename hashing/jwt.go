package hashing

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaim struct {
	Id    string `json:"id"`
	User  string `json:"user"`
	Admin bool   `json:"role"`
	jwt.RegisteredClaims
}

func CreateNewAuthToken(id string, email string, isAdmin bool) (string, error) {
	claims := AuthClaim{
		Id:    id,
		User:  email,
		Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "improve-yourself.com",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey, exists := os.LookupEnv("SECRET_KEY")
	if !exists {
		return "", errors.New("No secret key found")
	}

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("Error signing token")
	}

	return signedToken, nil
}

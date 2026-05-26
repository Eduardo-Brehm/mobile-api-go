package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// generateToken generates a JWT token with the userID and expiration time, and returns the token string
func GenerateToken(userID string) (string, error) {
	// get the secret key from the environment variable
	secret := os.Getenv("JWT_SECRET")

	// create a new JWT token with the userID and expiration time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,                                // Subject (user ID)
		"iat": time.Now().Unix(),                     // Issued at
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expiration (24 hours)
	})

	// sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies a JWT token and returns the userID if the token is valid, otherwise returns an error
func VerifyToken(tokenString string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	// extract the userID from the token claims
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)

	return userID, nil
}

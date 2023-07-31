package utils

import (
	"fmt"
	"time"

	. "clokify/config"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	secretKey := []byte(EnvConfig().JwtSecret)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseJWTToken(signedToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		// Return the secret key used for signing
		return []byte(EnvConfig().JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid and not expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

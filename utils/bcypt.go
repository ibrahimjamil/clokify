package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err == nil {
		return string(hashedPassword), nil
	}

	return "", err
}

func MatchPassword(hashedDBPassword string, userPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedDBPassword), []byte(userPassword))
}

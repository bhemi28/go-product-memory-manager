package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyPassword(dbPassword string, enteredPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(enteredPassword)); err != nil {
		log.Print("the password did not match.")
		return false
	}
	return true
}

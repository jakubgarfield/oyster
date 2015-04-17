package repository

import (
	"math/rand"
	"time"
)

var (
	PasswordCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+{}|:<>?,./;'[]-=1234567890")
)

func GeneratePassword(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	password := make([]rune, length)
	for i := range password {
		password[i] = PasswordCharacters[rand.Intn(len(PasswordCharacters))]
	}
	return string(password)
}

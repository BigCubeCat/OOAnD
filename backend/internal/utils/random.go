package utils

import (
	"time"

	"golang.org/x/exp/rand"

	"backend/internal/config"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomPassword() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	b := make([]rune, config.GetPasswordLen())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

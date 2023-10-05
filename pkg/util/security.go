package util

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(str string) string {
	hashedStr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	PanicIfNeeded(err)
	return string(hashedStr)
}

func VerifyHash(hashedStr, candidateStr string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(candidateStr)); err == nil {
		return true
	} else {
		return false
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

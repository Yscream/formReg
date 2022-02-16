package encryption

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 16
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func HashPassword(salt, password string) (string, error) {
	combination := salt + password
	bytes, err := bcrypt.GenerateFromPassword([]byte(combination), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("wrong password")
	}
	return string(bytes), err
}

func CheckPasswordAndHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

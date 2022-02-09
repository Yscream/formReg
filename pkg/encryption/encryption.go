package encryption

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func GenerateRandomString(secret []byte) []byte {
	const saltSize = 16
	buf := make([]byte, saltSize, saltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		fmt.Printf("random read failed: %v", err)
		os.Exit(1)
	}

	hash := sha1.New()
	hash.Write(buf)
	hash.Write(secret)
	return hash.Sum(buf)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("wrong password")
	}
	return string(bytes), err
}

func CheckPasswordAndHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

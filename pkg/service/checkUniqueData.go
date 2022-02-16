package service

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/DB"
	"github.com/Yscream/go-form-reg/pkg/encryption"
)

func CheckEmail(email string) bool {
	db := DB.GetConnection()

	defer db.Close()

	var existEmail string
	err := db.Get(&existEmail, "SELECT email FROM users_data WHERE email=$1", email)

	if err != nil {
		return false
	}

	if email != existEmail {
		return false
	}
	return true
}

func CheckPass(email, password string) error {
	db := DB.GetConnection()
	defer db.Close()

	var id int
	err := db.Get(&id, "SELECT id FROM users_data WHERE email=$1", email)
	if err != nil {
		return fmt.Errorf("id doesn't exist")
	}

	var salt string
	err = db.Get(&salt, "SELECT salt FROM credentials WHERE users_id=$1", id)
	if err != nil {
		return fmt.Errorf("salt doesn't exist")
	}

	var hash string
	err = db.Get(&hash, "SELECT hash FROM credentials WHERE users_id=$1", id)
	if err != nil {
		return fmt.Errorf("hash doesn't exist")
	}

	var combination = salt + password

	compare := encryption.CheckPasswordAndHash(combination, hash)

	return compare
}

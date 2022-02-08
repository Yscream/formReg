package validators

import (
	"database/sql"
	"fmt"
	"regexp"

	"github.com/Yscream/go-form-reg/configs"
	"github.com/Yscream/go-form-reg/pkg/encryption"
	_ "github.com/lib/pq"
)

func FieldLen(min, max int, name string) bool {
	if len(name) < min || len(name) > max {
		return false
	}
	return true
}

func Email(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(email) {
		return false
	}
	return true
}

func CheckEmail(email string) bool {
	conn := configs.GetConfig()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var existEmail string
	err = db.QueryRow("SELECT email FROM users_data WHERE email=$1", email).Scan(&existEmail)

	if err != nil {
		return false
	}

	if email != existEmail {
		return false
	}
	return true
}

func CheckPass(email, password string) error {
	conn := configs.GetConfig()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var id int
	err = db.QueryRow("SELECT id FROM users_data WHERE email=$1", email).Scan(&id)
	fmt.Println(id)

	var salt string
	err = db.QueryRow("SELECT salt FROM credentials WHERE users_id=$1", id).Scan(&salt)
	fmt.Println(salt)

	var hash string
	err = db.QueryRow("SELECT hash FROM credentials WHERE users_id=$1", id).Scan(&hash)
	fmt.Println(hash)

	var combination = salt + password
	fmt.Println(combination)

	compare := encryption.CheckPasswordAndHash(combination, hash)
	fmt.Println(compare == nil)
	return compare
}

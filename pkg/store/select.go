package store

import (
	"fmt"

	"github.com/Yscream/go-form-reg/configs"
	"github.com/jmoiron/sqlx"
)

func GetUser(email string) string {
	conn := configs.GetConfig()
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var name string
	var lname string
	err = db.QueryRow("SELECT fname, lname FROM users_data WHERE email=$1", email).Scan(&name, &lname)
	if err != nil {
		fmt.Println("id doesn't exist")
	}

	user := fmt.Sprintf("%s %s", name, lname)
	fmt.Println(user)
	return user
}

func GetId(email string) int {
	conn := configs.GetConfig()
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		panic(err)
	}

	var id int
	err = db.Get(&id, "SELECT id FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Println("id doesn't exist")
	}
	fmt.Println(id)

	return id
}

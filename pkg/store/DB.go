package store

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/Yscream/go-form-reg/configs"
	"github.com/Yscream/go-form-reg/pkg/encryption"
	"github.com/Yscream/go-form-reg/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func SaveData(user *models.User) {
	conn := configs.GetConfig()
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Fatalln(err)
	}

	salt := encryption.GenerateRandomString([]byte(user.Password))
	saltToString := base64.StdEncoding.EncodeToString(salt)
	combination := saltToString + user.Password
	hash, _ := encryption.HashPassword(combination)

	db.MustExec("INSERT INTO users_data (fname,lname,email)  VALUES($1, $2, $3)", user.Name, user.LastName, user.Email)
	db.MustExec("INSERT INTO credentials (users_id, salt, hash)  VALUES($1, $2, $3)", GetId(user.Email), saltToString, hash)

	fmt.Println("connect to server...")
}

func DeleteData(token string) {
	conn := configs.GetConfig()
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		panic(err)
	}

	db.MustExec("DELETE FROM tokens WHERE token=$1", token)
}

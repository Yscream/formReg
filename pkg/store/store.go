package store

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Yscream/formReg/pkg/encryption"
	"github.com/Yscream/formReg/pkg/models"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

func GetId(email string) int {
	conn := GetConfig()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var id int
	err = db.QueryRow("SELECT id FROM users_data WHERE email=$1", email).Scan(&id)
	if err != nil {
		fmt.Println("id doesn't exist")
	}
	fmt.Println(id)

	return id
}

func GetConfig() string {

	yfile, err := ioutil.ReadFile("./configs/config.yml")

	if err != nil {
		log.Fatal(err)
	}

	c := models.Config{}

	err = yaml.Unmarshal(yfile, &c)

	fmt.Println(c)

	if err != nil {

		log.Fatal(err)
	}
	result := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.Username, c.Password, c.Host, c.Port, c.Dbname, c.Sslmode)
	return result
}

func SaveData(user *models.User) {
	userdata := models.User{
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		Password: user.Password,
	}
	conn := GetConfig()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users_data (fname,lname,email)  VALUES($1, $2, $3)", &userdata.Name, &userdata.LastName, &userdata.Email)
	if err != nil {
		panic(err)
	}

	defer insert.Close()

	salt := encryption.GenerateRandomString([]byte(userdata.Password))
	saltToString := base64.StdEncoding.EncodeToString(salt)
	combination := saltToString + userdata.Password
	fmt.Printf("Salt : %x \n", salt)
	fmt.Println(saltToString)
	fmt.Println(combination)

	hash, _ := encryption.HashPassword(combination)
	fmt.Println(hash)

	credentials, err := db.Query("INSERT INTO credentials (users_id, salt, hash)  VALUES($1, $2, $3)", GetId(userdata.Email), saltToString, hash)
	if err != nil {
		panic(err)
	}

	defer credentials.Close()

	fmt.Println("connect to server...")
}

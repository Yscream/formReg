package DB

import "fmt"

func GetName(email string) string {
	db := GetConnection()

	var name string
	db.Get(&name, "SELECT fname FROM users_data WHERE email=$1", email)
	if db != nil {
		fmt.Println("smth wrong with fname, 10 string")
	}
	return name
}

func GetLname(email string) string {
	db := GetConnection()

	var lname string
	db.Get(&lname, "SELECT lname FROM users_data WHERE email=$1", email)
	if db != nil {
		fmt.Println("smth wrong with lname, 22 string")
	}
	return lname
}

func GetId(email string) int {
	db := GetConnection()
	defer db.Close()

	var id int
	db.Get(&id, "SELECT id FROM users_data WHERE email=$1", email)
	return id
}

func GetIdViaToken(token string) int {
	db := GetConnection()
	defer db.Close()

	var id int
	db.Get(&id, "SELECT id FROM tokens WHERE email=$1", token)
	return id
}

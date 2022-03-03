package DB

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/models"
)

func (m *DataBase) GetUser(email string) (string, string) {
	var name, lname string
	err := m.DBmodel.Get(&name, "SELECT fname FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = m.DBmodel.Get(&lname, "SELECT lname FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Println(err.Error())
	}

	return name, lname
}

func (m *DataBase) GetEmail(email string) error {
	var existEmail string
	err := m.DBmodel.Get(&existEmail, "SELECT email FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Println("Incorrect email")
	}
	return err
}

func (m *DataBase) GetId(email string) (int, error) {
	var id int
	err := m.DBmodel.Get(&id, "SELECT id FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Printf("%v doesn't exist", err)
	}
	return id, err
}

func (m *DataBase) GetSaltAndHash(id int) (string, string) {
	var salt, hash string
	err := m.DBmodel.Get(&salt, "SELECT salt FROM credentials WHERE users_id=$1", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = m.DBmodel.Get(&hash, "SELECT hash FROM credentials WHERE users_id=$1", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return salt, hash
}

func (m *DataBase) SaveData(salt, hash string, user *models.User) error {
	result, err := m.DBmodel.Queryx("INSERT INTO users_data (fname,lname,email)  VALUES($1, $2, $3) RETURNING id", user.Name, user.LastName, user.Email)
	if err != nil {
		fmt.Printf("smth wrong in requests, 57 string, %v", err)
		return err
	}
	for result.Next() {
		err = result.Scan(&user.ID)
		if err != nil {
			fmt.Println("failed to get id")
			return err
		}
	}
	m.DBmodel.Queryx("INSERT INTO credentials (users_id, salt, hash)  VALUES($1, $2, $3)", user.ID, salt, hash)

	return nil
}

func (m *DataBase) DeleteToken(token string) {
	m.DBmodel.Queryx("DELETE FROM tokens WHERE token=$1", token)

}

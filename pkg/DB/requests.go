package DB

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/models"
)

func (m *Connection) GetUser(email string) (string, string, error) {
	var name, lname string
	err := m.DBmodel.Get(&name, "SELECT fname FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = m.DBmodel.Get(&lname, "SELECT lname FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Println(err.Error())
	}

	return name, lname, err
}

func (m *Connection) GetEmail(email string) error {
	var existEmail string
	err := m.DBmodel.Get(&existEmail, "SELECT email FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Println("Incorrect email")
	}
	return err
}

func (m *Connection) GetId(email string) (int, error) {
	var id int
	err := m.DBmodel.Get(&id, "SELECT id FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Printf("%v doesn't exist", err)
	}
	return id, err
}

func (m *Connection) GetSaltAndHash(id int) (string, string, error) {
	var salt, hash string
	err := m.DBmodel.Get(&salt, "SELECT salt FROM credentials WHERE users_id=$1", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = m.DBmodel.Get(&hash, "SELECT hash FROM credentials WHERE users_id=$1", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return salt, hash, err
}

func (m *Connection) InsertUser(user *models.User) error {
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
	return nil
}

func (m *Connection) InsertPassword(id int, salt, hash string) error {
	_, err := m.DBmodel.Queryx("INSERT INTO credentials (users_id, salt, hash)  VALUES($1, $2, $3)", id, salt, hash)
	if err != nil {
		fmt.Print("cannot insert password", err)
		return err
	}
	return nil
}

func (m *Connection) InsertToken(id int, token string) error {
	_, err := m.DBmodel.Exec("INSERT INTO tokens (users_id, token)  VALUES($1, $2)", id, token)
	if err != nil {
		fmt.Println("cannot delete token", err)
		return err
	}
	return nil
}

func (m *Connection) DeleteToken(token string) error {
	_, err := m.DBmodel.Exec("DELETE FROM tokens WHERE token=$1", token)
	if err != nil {
		fmt.Println("cannot delete token", err)
		return err
	}
	return nil
}

func (m *Connection) SelectToken(id int) (string, error) {
	var token string
	err := m.DBmodel.Get(&token, "SELECT token FROM tokens WHERE users_id=$1", id)
	if err != nil {
		fmt.Println("cannot select token", err)
		return "", err
	}
	return token, err
}

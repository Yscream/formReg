package postgresql

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/models"
)

func (m *Repository) GetUser(email string) (string, string, error) {
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

func (m *Repository) GetEmail(email string) (string, error) {
	var existEmail string
	err := m.DBmodel.Get(&existEmail, "SELECT email FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Println("Incorrect email")
	}
	return existEmail, err
}

func (m *Repository) GetEmailByToken(token string) (string, error) {
	var email string
	err := m.DBmodel.Get(&email, "SELECT u.email FROM users_data u, tokens t WHERE t.id=$1", token)
	if err != nil {
		fmt.Println(err)
	}
	return email, nil
}

func (m *Repository) GetId(email string) (int, error) {
	var id int
	err := m.DBmodel.Get(&id, "SELECT id FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Printf("%v doesn't exist", err)
	}
	return id, err
}

func (m *Repository) GetCredentials(id int) (models.Credentials, error) {
	cred := models.Credentials{}

	err := m.DBmodel.Get(&cred, "SELECT * FROM credentials WHERE users_id=$1", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	// err = m.DBmodel.Get(&hash, "SELECT hash FROM credentials WHERE users_id=$1", id)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	return cred, err
}

func (m *Repository) InsertUser(user *models.User) error {
	result, err := m.DBmodel.Queryx("INSERT INTO users_data (fname,lname,email)  VALUES($1, $2, $3) RETURNING id", user.Name, user.LastName, user.Email)
	if err != nil {
		fmt.Printf("cannot insert user, %v", err.Error())
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

func (m *Repository) InsertPassword(id int, salt, hash string) error {
	_, err := m.DBmodel.Queryx("INSERT INTO credentials (users_id, salt, hash)  VALUES($1, $2, $3)", id, salt, hash)
	if err != nil {
		fmt.Print("cannot insert password", err)
		return err
	}
	return nil
}

func (m *Repository) InsertToken(id int, token string) error {
	_, err := m.DBmodel.Queryx("INSERT INTO tokens (users_id, token)  VALUES($1, $2)", id, token)
	if err != nil {
		fmt.Println("cannot delete token", err)
		return err
	}
	return nil
}

func (m *Repository) SelectToken(id int) (string, error) {
	var token string
	err := m.DBmodel.Get(&token, "SELECT token FROM tokens WHERE users_id=$1", id)
	if err != nil {
		fmt.Println("cannot select token", err)
		return "", err
	}
	return token, err
}

func (m *Repository) DeleteToken(token string) error {
	_, err := m.DBmodel.Queryx("DELETE FROM tokens WHERE token=$1", token)
	if err != nil {
		fmt.Println("cannot delete token", err)
		return err
	}
	return nil
}

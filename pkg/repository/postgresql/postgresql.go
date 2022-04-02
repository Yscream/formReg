package postgresql

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/models"
)

func (m *Repository) GetUser(email string) (models.Person, error) {
	user := models.Person{}
	err := m.DBmodel.Get(&user, "SELECT * FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Println(err.Error())
	}
	return user, err
}

func (m *Repository) GetEmail(email string) (string, error) {
	var existEmail string
	err := m.DBmodel.Get(&existEmail, "SELECT email FROM users_data WHERE email=$1", email)
	if err != nil {
		fmt.Printf("Cannot find email %s", err)
		return "", err
	}

	return existEmail, nil
}

//testing
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
	return cred, err
}

func (m *Repository) InsertUser(user *models.User) error {
	insert, err := m.DBmodel.Queryx("INSERT INTO users_data (fname,lname,email)  VALUES($1, $2, $3) RETURNING id", user.Name, user.LastName, user.Email)
	if err != nil {
		fmt.Printf("cannot insert user, %v", err.Error())
		return err
	}
	defer insert.Close()

	for insert.Next() {
		err = insert.Scan(&user.ID)
		if err != nil {
			fmt.Println("failed to get id")
			return err
		}
	}
	return nil
}

func (m *Repository) InsertCredentials(cred *models.Credentials) error {
	insert, err := m.DBmodel.Queryx("INSERT INTO credentials (users_id, salt, hash)  VALUES($1, $2, $3)", cred.ID, cred.Salt, cred.Hash)
	if err != nil {
		fmt.Print("cannot insert password", err)
		return err
	}
	defer insert.Close()

	return nil
}

func (m *Repository) InsertToken(at *models.AccessToken) error {
	insert, err := m.DBmodel.Queryx("INSERT INTO tokens (users_id, token)  VALUES($1, $2)", at.ID, at.Token)
	if err != nil {
		fmt.Println("cannot delete token", err)
		return err
	}
	defer insert.Close()

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
	delete, err := m.DBmodel.Queryx("DELETE FROM tokens WHERE token=$1", token)
	if err != nil {
		fmt.Println("cannot delete token", err)
		return err
	}
	defer delete.Close()

	return nil
}

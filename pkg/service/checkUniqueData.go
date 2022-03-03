package service

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/encryption"
)

func CheckEmail(email string, app *Application) error {
	existEmail := app.data.GetEmail(email)
	fmt.Println(existEmail)
	if existEmail != nil {
		return existEmail
	}
	return nil
}

func CheckPass(email, password string, app *Application) error {
	id, err := app.data.GetId(email)
	if err != nil {
		fmt.Println(err)
	}

	salt, hash := app.data.GetSaltAndHash(id)

	fmt.Println(salt, hash)

	var combination = salt + password

	compare := encryption.CheckPasswordAndHash(combination, hash)
	fmt.Println(compare)

	return compare
}

package service

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/encryption"
)

func CheckEmail(email string, app *Application) error {
	_, ok := app.data.GetEmail(email)
	fmt.Println(ok)
	if ok != nil {
		return ok
	}
	return nil
}

func CheckPass(email, password string, app *Application) error {
	id, ok := app.data.GetId(email)
	if ok != nil {
		fmt.Println(ok)
	}

	salt, hash, err := app.data.GetSaltAndHash(id)

	if err != nil {
		fmt.Print(err.Error())
	}

	var combination = salt + password

	compare := encryption.CheckPasswordAndHash(combination, hash)
	fmt.Println(compare)

	return compare
}

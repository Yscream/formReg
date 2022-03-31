package service

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/encryption"
)

func CheckPass(email, password string, app *Application) error {
	id, ok := app.data.GetId(email)
	if ok != nil {
		fmt.Println(ok)
	}

	cred, err := app.data.GetCredentials(id)

	if err != nil {
		fmt.Print(err.Error())
	}

	var combination = cred.Salt + password

	compare := encryption.CheckPasswordAndHash(combination, cred.Hash)
	fmt.Println(compare)

	return compare
}

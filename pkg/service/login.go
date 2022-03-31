package service

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/models"
)

var Email string

func Login(user *models.LoginUser, app *Application) []models.TypeOfErrors {
	errors := make([]models.TypeOfErrors, 0)

	checkEmail, _ := app.data.GetEmail(user.Email)
	if checkEmail != user.Email {
		fmt.Println("from login", checkEmail)
		errors = append(errors, models.TypeOfErrors{
			FieldName:  "Email",
			MessageErr: "Incorrect email address",
		})
	}

	checkPass := CheckPass(user.Email, user.Password, app)
	if checkPass != nil {
		errors = append(errors, models.TypeOfErrors{
			FieldName:  "Password",
			MessageErr: "incorrect password",
		})
	}
	return errors
}

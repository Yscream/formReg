package service

import (
	"github.com/Yscream/go-form-reg/pkg/models"
)

var Email string

func Login(user *models.LoginUser, app *Application) []models.TypeOfErrors {
	errors := make([]models.TypeOfErrors, 0)

	_, err := app.data.GetEmail(user.Email)
	if err != nil {
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

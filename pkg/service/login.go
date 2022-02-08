package service

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/models"
	"github.com/Yscream/go-form-reg/pkg/validators"
)

func Login(user *models.LoginUser) []models.TypeOfErrors {
	errors := make([]models.TypeOfErrors, 0)

	checkEmail := validators.CheckEmail(user.Email)
	if !checkEmail {
		fmt.Println("from login", checkEmail)
		errors = append(errors, models.TypeOfErrors{
			FieldName:  "Email",
			MessageErr: "Incorrect email address",
		})
	}

	checkPass := validators.CheckPass(user.Email, user.Password)
	if checkPass != nil {
		errors = append(errors, models.TypeOfErrors{
			FieldName:  "Password",
			MessageErr: "incorrect password",
		})
	}
	return errors
}

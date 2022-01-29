package service

import (
	"fmt"

	"github.com/Yscream/formReg/pkg/models"
	"github.com/Yscream/formReg/pkg/validator"
)

func Login(user *models.LoginUser) []models.TypeOfErrors {
	errors := make([]models.TypeOfErrors, 0)

	checkEmail := validator.CheckEmail(user.Email)
	if !checkEmail {
		fmt.Println("from login", checkEmail)
		errors = append(errors, models.TypeOfErrors{
			FieldName:  "Email",
			MessageErr: "Incorrect email address",
		})
	}

	checkPass := validator.CheckPass(user.Email, user.Password)
	if checkPass != nil {
		errors = append(errors, models.TypeOfErrors{
			FieldName:  "Password",
			MessageErr: "incorrect password",
		})
	}
	return errors
}

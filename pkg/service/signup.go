package service

import (
	"github.com/Yscream/go-form-reg/pkg/models"
	"github.com/Yscream/go-form-reg/pkg/validation"
)

func Signup(user *models.User) []models.TypeOfErrors {
	errors := make([]models.TypeOfErrors, 0)
	validName := validation.FieldLen(2, 255, user.Name)
	if !validName {
		errors = append(errors, models.TypeOfErrors{
			FieldName:  "Name",
			MessageErr: "First name must contains at least 2 and no more than 255 symbols",
		})
	}
	validLastName := validation.FieldLen(2, 255, user.LastName)
	if !validLastName {
		errors = append(errors, models.TypeOfErrors{
			FieldName:  "LastName",
			MessageErr: "Last name must contains at least 2 and no more than 255 symbols",
		})
	}
	validEmail := validation.Email(user.Email)
	if !validEmail {
		errors = append(errors, models.TypeOfErrors{
			FieldName:  "Email",
			MessageErr: "Incorrect email address",
		})
	}
	CheckEmail := CheckEmail(user.Email)

	if CheckEmail {
		errors = append(errors, models.TypeOfErrors{
			FieldName:  "Email",
			MessageErr: "Email address already registered",
		})
	}

	validPassword := validation.FieldLen(8, 64, user.Password)
	if !validPassword {
		errors = append(errors, models.TypeOfErrors{
			FieldName:  "Password",
			MessageErr: "Password must contains at least 8 symbols and no more than 64 symbols",
		})
	}
	return errors
}

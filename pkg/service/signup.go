package service

import (
	"encoding/base64"

	"github.com/Yscream/go-form-reg/pkg/encryption"
	"github.com/Yscream/go-form-reg/pkg/models"
	"github.com/Yscream/go-form-reg/pkg/validation"
)

func (app *Application) HashingPassword(user *models.User) error {
	salt := encryption.GenerateRandomString([]byte(user.Password))
	hash, _ := encryption.HashPassword(base64.StdEncoding.EncodeToString(salt), user.Password)
	err := app.data.InsertUser(user)
	if err != nil {
		return err
	}

	err = app.data.InsertPassword(user.ID, base64.StdEncoding.EncodeToString(salt), hash)
	if err != nil {
		return err
	}

	return nil
}

func Signup(user *models.User, app *Application) []models.TypeOfErrors {
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
	CheckEmail := CheckEmail(user.Email, app)
	if CheckEmail == nil {
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

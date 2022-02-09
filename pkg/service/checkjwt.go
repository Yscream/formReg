package service

import (
	"github.com/Yscream/go-form-reg/pkg/DB"
	"github.com/Yscream/go-form-reg/pkg/JWT"
	"github.com/Yscream/go-form-reg/pkg/models"
)

func CheckJWT(email, head string, hmacSecret []byte) []models.Person {
	errors := make([]models.Person, 0)
	check := JWT.ParseJWT(head, hmacSecret)
	if check != nil {
		errors = append(errors, models.Person{
			Tokenerr: check.Error(),
		})
		return errors
	}
	errors = append(errors, models.Person{
		Name:  DB.GetName(email),
		Lname: DB.GetLname(email),
		Email: email,
	})
	return errors
}

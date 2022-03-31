package service

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/JWT"
	"github.com/Yscream/go-form-reg/pkg/models"

	"github.com/golang-jwt/jwt"
)

func (app *Application) CheckJWT(email, token string, hmacSecret []byte) []models.Person {
	person, err := app.data.GetUser(email)
	if err != nil {
		fmt.Println("cannot take user:", err)
	}
	errors := make([]models.Person, 0)
	check := ParseJWT(token, hmacSecret)
	if check != nil {
		app.data.DeleteToken(token)
		errors = append(errors, models.Person{
			Token: check.Error(),
		})
		return errors
	}
	errors = append(errors, models.Person{
		ID:    person.ID,
		Name:  person.Name,
		Lname: person.Lname,
		Email: email,
	})
	return errors
}

func (app *Application) SaveToken(user *models.LoginUser) {
	id, err := app.data.GetId(user.Email)
	if err != nil {
		fmt.Println(err)
	}
	person, err := app.data.GetUser(user.Email)
	if err != nil {
		fmt.Println("cannot take user:", err)
	}

	token, err := JWT.NewJWT(user.Email, person.Name, person.Lname, id)
	if err != nil {
		fmt.Println(err)
	}

	res := app.data.InsertToken(id, token)
	if res != nil {
		fmt.Println(res)
	}
}

func (app *Application) DeleteToken(token string) {
	res := app.data.DeleteToken(token)
	if res != nil {
		fmt.Println(res)
	}
}

func (app *Application) SelectToken(user *models.LoginUser) string {
	id, err := app.data.GetId(user.Email)
	if err != nil {
		fmt.Println(err)
	}
	token, err := app.data.SelectToken(id)
	if err != nil {
		fmt.Println(err.Error())
	}

	return token
}

func ParseJWT(tokenStr string, hmacSecret []byte) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("ololo")
		return nil
	} else {
		fmt.Println(err)
		return err
	}
}

package service

import (
	"fmt"

	"github.com/Yscream/go-form-reg/pkg/JWT"
	"github.com/Yscream/go-form-reg/pkg/models"

	"github.com/golang-jwt/jwt"
)

func (app *Application) CheckJWT(email, token string, hmacSecret []byte) []models.Person {
	name, lname := app.data.GetUser(email)
	errors := make([]models.Person, 0)
	check := ParseJWT(token, hmacSecret)
	if check != nil {
		// app.data.DeleteToken(token)
		errors = append(errors, models.Person{
			Tokenerr: check.Error(),
		})
		return errors
	}
	errors = append(errors, models.Person{
		Name:  name,
		Lname: lname,
		Email: email,
	})
	return errors
}

func (app *Application) SaveToken(user *models.LoginUser) {
	id, err := app.data.GetId(user.Email)
	if err != nil {
		fmt.Println(err)
	}
	name, lname := app.data.GetUser(user.Email)

	token, err := JWT.NewJWT(user.Email, name, lname, id)
	if err != nil {
		fmt.Println(err)
	}
	app.data.DBmodel.Exec("INSERT INTO tokens (users_id, token)  VALUES($1, $2)", id, token)
}

func (app *Application) DeleteToken(token string) {
	app.data.DeleteToken(token)
}

func (app *Application) SendToken(user *models.LoginUser) string {
	id, err := app.data.GetId(user.Email)
	if err != nil {
		fmt.Println(err)
	}
	var token string
	err = app.data.DBmodel.Get(&token, "SELECT token FROM tokens WHERE users_id=$1", id)
	if err != nil {
		fmt.Println(err)
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

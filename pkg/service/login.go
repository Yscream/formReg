package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/models"
)

var Email string

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	login := models.LoginUser{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal([]byte(body), &login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Email = login.Email

	errors := Login(&login, app)

	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		marshalBytes, err := json.Marshal(&errors)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(marshalBytes)
		return
	}
	app.SaveToken(&login)
	token, err := json.Marshal(app.SendToken(&login))
	if err != nil {
		return
	}
	w.Write(token)
}

func Login(user *models.LoginUser, app *Application) []models.TypeOfErrors {
	errors := make([]models.TypeOfErrors, 0)

	checkEmail := CheckEmail(user.Email, app)
	if checkEmail != nil {
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

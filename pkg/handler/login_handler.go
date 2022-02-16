package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/JWT"
	"github.com/Yscream/go-form-reg/pkg/models"
	"github.com/Yscream/go-form-reg/pkg/service"
)

var Email string

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println(Email)
	errors := service.Login(&login)

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
	JWT.SaveToken(&login)
	token, err := json.Marshal(JWT.SendToken(&login))
	if err != nil {
		return
	}
	w.Write(token)
}

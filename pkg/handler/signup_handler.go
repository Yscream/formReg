package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/DB"
	"github.com/Yscream/go-form-reg/pkg/models"
	"github.com/Yscream/go-form-reg/pkg/service"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("Form :%+v;\n", user)

	errors := service.Signup(&user)

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
	DB.SaveData(&user)
	w.Write([]byte("[]"))
}

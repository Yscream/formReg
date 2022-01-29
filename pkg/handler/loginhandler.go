package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Yscream/formReg/pkg/models"
	"github.com/Yscream/formReg/pkg/service"
	"github.com/Yscream/formReg/pkg/store"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	login := models.LoginUser{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal([]byte(body), &login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("Form :%+v;\n", login)

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

	w.Write([]byte("[]"))
	store.SaveToken(&login)
}

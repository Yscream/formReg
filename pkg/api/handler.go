package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/JWT"
	"github.com/Yscream/go-form-reg/pkg/models"
	"github.com/Yscream/go-form-reg/pkg/service"
)

type ServiceHandler struct {
	service *service.Application
}

func NewHandler(service *service.Application) *ServiceHandler {
	return &ServiceHandler{
		service: service,
	}
}

func (handler *ServiceHandler) NewSignupHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("38", err)
	errors := service.Signup(&user, handler.service)
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
	fmt.Println("50", err)
	err = handler.service.InsertUserData(&user)
	fmt.Println("52", err)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[]"))
}

var Email string

func (handler *ServiceHandler) NewLogInHandler(w http.ResponseWriter, r *http.Request) {
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
	errors := service.Login(&login, handler.service)
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
	handler.service.SaveToken(&login)
	selectToken := handler.service.SelectToken(&login)
	token, err := json.Marshal(selectToken)
	if err != nil {
		return
	}
	w.Write(token)
}

func (handler *ServiceHandler) NewLogOutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	handler.service.DeleteToken(token)
}

func (handler *ServiceHandler) ShowProfile(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	profile := handler.service.CheckJWT(Email, token, JWT.HmacSampleSecret)

	marshal, err := json.Marshal(&profile)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(marshal)
}

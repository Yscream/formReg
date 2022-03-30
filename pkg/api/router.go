package api

import (
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/service"
	"github.com/gorilla/mux"
)

func NewRouter(service *service.Application) *mux.Router {
	router := mux.NewRouter()
	handler := NewHandler(service)
	router.HandleFunc("/user", handler.NewSignupHandler)
	router.HandleFunc("/log", handler.NewLogInHandler)
	router.HandleFunc("/log_out", handler.NewLogOutHandler)
	router.HandleFunc("/token", handler.ShowProfile)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../assets/")))
	return router
}

package router

import (
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/handler"
	"github.com/Yscream/go-form-reg/pkg/service"
	"github.com/gorilla/mux"
)

func NewRouter(service *service.Application) *mux.Router {
	router := mux.NewRouter()
	h := handler.NewHandler(service)
	router.HandleFunc("/user", h.NewSignupHandler)
	router.HandleFunc("/log", h.NewLogInHandler)
	router.HandleFunc("/log_out", h.NewLogOutHandler)
	router.HandleFunc("/token", h.ShowProfile)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../assets/")))
	return router
}

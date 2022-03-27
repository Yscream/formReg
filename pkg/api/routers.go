package api

import (
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/service"
	"github.com/gorilla/mux"
)

// func HandleHTML(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		http.ServeFile(w, r, "../assets/index.html")
// 		return
// 	}
// 	http.ServeFile(w, r, "../assets"+r.URL.Path)
// }

func NewRouters(service *service.Application) *mux.Router {
	router := mux.NewRouter()
	handler := NewHandler(service)
	router.HandleFunc("/user", handler.NewSignupHandler)
	router.HandleFunc("/log", handler.NewLogInHandler)
	router.HandleFunc("/log_out", handler.NewLogOutHandler)
	router.HandleFunc("/token", handler.ShowProfile)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
	return router
}

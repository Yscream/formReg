package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Yscream/formReg/pkg/handler"
	_ "github.com/lib/pq"
)

func HandleHTML(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "./assets/index.html")
}

func main() {
	http.HandleFunc("/", HandleHTML)
	http.HandleFunc("/user", handler.SignupHandler)
	http.HandleFunc("/log", handler.LoginHandler)
	http.Handle("/link.html", http.FileServer(http.Dir("./assets")))
	http.Handle("/after_log.html", http.FileServer(http.Dir("./assets")))
	http.Handle("/log.html", http.FileServer(http.Dir("./assets")))
	http.Handle("/index.js", http.FileServer(http.Dir("./assets")))
	http.Handle("/login.js", http.FileServer(http.Dir("./assets")))
	http.Handle("/style.css", http.FileServer(http.Dir("./assets")))
	fmt.Printf("Starting server for testing HTTP POST... PORT: 8033\n")
	if err := http.ListenAndServe(":8033", nil); err != nil {
		log.Fatal(err)
	}
}

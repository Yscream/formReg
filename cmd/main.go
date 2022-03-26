package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Yscream/go-form-reg/configs"
	"github.com/Yscream/go-form-reg/pkg/handler"
	"github.com/Yscream/go-form-reg/pkg/repository/postgresql"
	"github.com/Yscream/go-form-reg/pkg/service"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
)

func HandleHTML(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "../assets/index.html")
		return
	}
	http.ServeFile(w, r, "../assets"+r.URL.Path)
}

func main() {
	conn, err := configs.InitConfig("../cmd/config.yml")
	if err != nil {
		log.Fatalf("cannot read config")
	}
	db, err := postgresql.OpenDB(conn)
	if err != nil {
		log.Fatal(err)
	}

	connection := service.NewConnection(db)

	http.HandleFunc("/user", handler.NewSignupHandler(connection))
	http.HandleFunc("/log", handler.NewLogInHandler(connection))
	http.HandleFunc("/log_out", handler.NewLogOutHandler(connection))
	http.HandleFunc("/token", handler.NewProfile(connection))
	http.HandleFunc("/", HandleHTML)
	fmt.Printf("Starting server for testing HTTP POST... PORT: 8033\n")
	if err := http.ListenAndServe("0.0.0.0:8033", nil); err != nil {
		log.Fatal(err)
	}
}

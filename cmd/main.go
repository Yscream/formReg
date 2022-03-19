package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Yscream/go-form-reg/configs"
	"github.com/Yscream/go-form-reg/pkg/handler"
	"github.com/Yscream/go-form-reg/pkg/service"
	"github.com/jmoiron/sqlx"
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
	db, err := OpenDB(conn)
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

func OpenDB(conn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", conn)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(95)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Yscream/go-form-reg/configs"
	"github.com/Yscream/go-form-reg/pkg/repository/postgresql"
	"github.com/Yscream/go-form-reg/pkg/router"
	"github.com/Yscream/go-form-reg/pkg/service"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := configs.InitConfig("../cmd/config.yml")
	if err != nil {
		log.Fatalf("cannot read config")
	}
	db, err := postgresql.OpenDB(conn)
	if err != nil {
		log.Fatal(err)
	}
	repository := service.NewConnection(db)

	fmt.Printf("Starting server for testing HTTP POST... PORT: 8033\n")
	log.Fatal(http.ListenAndServe(":8033", router.NewRouter(repository)))
}

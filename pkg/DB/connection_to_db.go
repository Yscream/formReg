package DB

import (
	"log"

	"github.com/Yscream/go-form-reg/configs"
	"github.com/jmoiron/sqlx"
)

func GetConnection() *sqlx.DB {
	conn := configs.GetConfig()
	db, err := sqlx.Open("postgres", conn)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(95)
	if err != nil {
		log.Println("m=GetConnection,msg=connection has failed", err)
	}
	return db
}

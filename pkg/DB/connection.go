package DB

import (
	"github.com/jmoiron/sqlx"
)

type DataBase struct {
	DBmodel *sqlx.DB
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

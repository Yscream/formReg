package DB

import (
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	DBmodel *sqlx.DB
}

func NewConnection(db *sqlx.DB) *Connection {
	return &Connection{
		DBmodel: db,
	}
}

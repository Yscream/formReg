package DB

import (
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:2201@localhost:6080/users?sslmode=disable"
)

var db *Connection

func TestMain(m *testing.M) {
	conn, err := sqlx.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	db = NewConnection(conn)

	os.Exit(m.Run())
}

package postgresql

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/pkg/errors"
)

var testURL string

func TestMain(m *testing.M) {
	var exitCode int
	defer os.Exit(exitCode)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.1",
		Env:        []string{"POSTGRES_USER=postgres", "POSTGRES_PASSWORD=2201", "POSTGRES_DB=users"},
		Name:       "repository_integration_tests",
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		var DB *sqlx.DB
		testURL = fmt.Sprintf("postgres://postgres:2201@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), "users")
		DB, err = sqlx.Connect("postgres", testURL)
		if err != nil {
			return err
		}
		return DB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	err = migrations("./migrations", testURL)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	exitCode = m.Run()
}

func getDB(t *testing.T, dbURL string) *Repository {
	t.Helper()

	db, err := OpenDB(dbURL)
	if err != nil {
		t.Fatalf("Could not connect to db, err: %s", err.Error())
	}

	return db
}

func migrations(migs, dbURL string) error {
	mgrt, err := migrate.New("file://"+migs, dbURL)
	if err != nil {
		return errors.Errorf("Could not open migration sources: %s, err: %s", migs, err.Error())
	}

	err = mgrt.Up()
	if err != nil {
		return errors.Errorf("Could not apply migrations: %s, err: %s", migs, err.Error())
	}
	return nil
}

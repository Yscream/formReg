package postgresql

import "github.com/jmoiron/sqlx"

type Repository struct {
	DBmodel *sqlx.DB
}

func OpenDB(conn string) (*Repository, error) {
	db, err := sqlx.Connect("postgres", conn)
	// db.SetMaxIdleConns(5)
	// db.SetMaxOpenConns(95)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &Repository{
		DBmodel: db,
	}, nil
}

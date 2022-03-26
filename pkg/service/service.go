package service

import (
	"github.com/Yscream/go-form-reg/pkg/repository/postgresql"
)

type Application struct {
	data *postgresql.Repository
}

func NewConnection(db *postgresql.Repository) *Application {
	return &Application{
		data: db,
	}
}

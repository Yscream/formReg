package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/DB"
	"github.com/Yscream/go-form-reg/pkg/JWT"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	data *DB.Connection
}

func NewConnection(db *sqlx.DB) *Application {
	return &Application{
		data: &DB.Connection{
			DBmodel: db,
		},
	}
}

func (app *Application) ShowProfile(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("Authorization")
	profile := app.CheckJWT(Email, head, JWT.HmacSampleSecret)

	marshal, err := json.Marshal(&profile)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(marshal)
}

package handler

import (
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/service"
)

var Email string

func NewLogInHandler(app *service.Application) http.HandlerFunc {
	return app.LoginHandler
}

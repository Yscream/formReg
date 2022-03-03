package handler

import (
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/service"
)

func NewProfile(app *service.Application) http.HandlerFunc {
	return app.ShowProfile
}

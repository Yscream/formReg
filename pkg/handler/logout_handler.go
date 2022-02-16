package handler

import (
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/JWT"
)

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("Authorization")
	JWT.DeleteToken(head)
}

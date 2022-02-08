package handler

import (
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/store"
)

func CheckToken(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("Authorization")
	store.ParseJWT(head, store.HmacSampleSecret)
}

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("Authorization")
	store.DeleteData(head)
}

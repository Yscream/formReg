package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/JWT"
	"github.com/Yscream/go-form-reg/pkg/service"
)

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("Authorization")
	JWT.DeleteData(head)
}

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("Authorization")
	fmt.Println(Email)
	profile := service.CheckJWT(Email, head, JWT.HmacSampleSecret)

	marshal, err := json.Marshal(&profile)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(marshal)
}

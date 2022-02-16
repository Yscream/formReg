package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/JWT"
	"github.com/Yscream/go-form-reg/pkg/service"
)

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("Authorization")
	profile := service.CheckJWT(Email, head, JWT.HmacSampleSecret)

	marshal, err := json.Marshal(&profile)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(marshal)
}

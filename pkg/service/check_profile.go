package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Yscream/go-form-reg/pkg/JWT"
)

func (app *Application) ShowProfile(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("Authorization")
	profile := app.CheckJWT(Email, head, JWT.HmacSampleSecret)

	marshal, err := json.Marshal(&profile)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(marshal)
}

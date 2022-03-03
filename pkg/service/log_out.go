package service

import "net/http"

func (app *Application) LogOutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	app.DeleteToken(token)
}

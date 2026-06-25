package app

import (
	"net/http"
)

func (app *Application) HelthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

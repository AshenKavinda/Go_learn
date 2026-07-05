package utils

import (
	"log"
	"net/http"
)

func WriteError(w http.ResponseWriter, status int, err error) error {
	type envelop struct {
		Error string `json:"error"`
	}
	log.Println(err)
	return WriteJSON(w, status, &envelop{Error: err.Error()})
}

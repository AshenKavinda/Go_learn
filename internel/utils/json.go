package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 //1_mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decorder := json.NewDecoder(r.Body)
	decorder.DisallowUnknownFields()

	return decorder.Decode(data)
}

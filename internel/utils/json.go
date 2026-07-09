package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	appError "github.com/ashenkavinda/go_social_app/internel/error"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 //1_mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decorder := json.NewDecoder(r.Body)
	decorder.DisallowUnknownFields()

	if err := decorder.Decode(data); err != nil {
		return err
	}
	return nil
}

func WriteError(w http.ResponseWriter, err error) {

	type envelope struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Details interface{} `json:"details"`
	}

	var appErr *appError.AppError

	if errors.As(err, &appErr) {

		if appErr.Err != nil {
			log.Println(appErr.Err)
		}

		WriteJSON(
			w,
			appErr.Status,
			&envelope{
				Code:    appErr.Code,
				Message: appErr.Message,
				Details: appErr.Details,
			},
		)

		return
	}

	// unknown error
	WriteJSON(
		w,
		http.StatusInternalServerError,
		&envelope{
			Code:    "INTERNAL_ERROR",
			Message: "Internal server error",
		},
	)
}

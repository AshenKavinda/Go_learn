package appError

import (
	"net/http"
)

type AppError struct {
	Code    string
	Message string
	Status  int
	Details interface{}
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NotFound(msg string) *AppError {
	return &AppError{
		Code:    "NOT_FOUND",
		Message: msg,
		Status:  http.StatusNotFound,
	}
}

func BadRequest(msg string, details ...interface{}) *AppError {

	var d interface{}
	if len(details) > 0 {
		d = details[0]
	}

	return &AppError{
		Code:    "BAD_REQUEST",
		Message: msg,
		Status:  http.StatusBadRequest,
		Details: d,
	}
}

func Internel(err error) *AppError {
	return &AppError{
		Code:    "INTERNEL_ERROR",
		Message: "Internal server error",
		Status:  http.StatusInternalServerError,
		Err:     err,
	}
}

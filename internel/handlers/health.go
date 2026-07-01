package handlers

import (
	"net/http"

	"github.com/ashenkavinda/go_social_app/internel/store"
)

type Handlers struct {
	Store store.Storage
}

func NewHandlers(store store.Storage) *Handlers {
	return &Handlers{
		Store: store,
	}
}

func (h *Handlers) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

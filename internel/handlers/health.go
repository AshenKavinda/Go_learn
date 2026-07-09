package handlers

import (
	"net/http"

	"github.com/ashenkavinda/go_social_app/internel/config"
	"github.com/ashenkavinda/go_social_app/internel/utils"
)

type HealthHandler struct {
	config config.AppConfig
}

func NewHealthHandler(config config.AppConfig) *HealthHandler {
	return &HealthHandler{
		config: config,
	}
}

func (h *HealthHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":  "ok",
		"env":     h.config.ENV,
		"version": h.config.Version,
	}

	utils.WriteJSON(w, 200, data)
}

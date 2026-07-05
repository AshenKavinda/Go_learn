package handlers

import (
	"net/http"

	"github.com/ashenkavinda/go_social_app/internel/dto/request"
	"github.com/ashenkavinda/go_social_app/internel/service"
	"github.com/ashenkavinda/go_social_app/internel/utils"
)

type PostHandler struct {
	PostService *service.PostService
}

func NewPostHandler(s *service.PostService) *PostHandler {
	return &PostHandler{
		PostService: s,
	}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	dto := &request.PostRequest{}
	ctx := r.Context()
	if err := utils.ReadJSON(w, r, dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	responce, err := h.PostService.Create(ctx, dto)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, responce)

}

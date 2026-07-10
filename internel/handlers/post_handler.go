package handlers

import (
	"log"
	"net/http"

	"github.com/ashenkavinda/go_social_app/internel/dto/request"
	appError "github.com/ashenkavinda/go_social_app/internel/error"
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
		utils.WriteError(w, appError.BadRequest("request validation error"))
		return
	}

	if errs := utils.ValidateStruct(dto); errs != nil {
		log.Println(errs)
		utils.WriteError(w, appError.BadRequest("validation error", errs))
		return
	}

	responce, err := h.PostService.Create(ctx, dto)
	if err != nil {
		utils.WriteError(w, appError.Internel(err))
		return
	}
	utils.WriteJSON(w, http.StatusCreated, responce)

}

func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := h.PostService.GetAll(r.Context())
	if err != nil {
		utils.WriteError(w, err)
	}

	utils.WriteJSON(w, http.StatusOK, posts)
}

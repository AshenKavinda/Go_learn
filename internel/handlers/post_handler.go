package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ashenkavinda/go_social_app/internel/dto/request"
	"github.com/ashenkavinda/go_social_app/internel/dto/response"
	appError "github.com/ashenkavinda/go_social_app/internel/error"
	"github.com/ashenkavinda/go_social_app/internel/service"
	"github.com/ashenkavinda/go_social_app/internel/utils"
	"github.com/go-chi/chi/v5"
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
	dto := &request.CreatePost{}
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
		return
	}

	utils.WriteJSON(w, http.StatusOK, posts)
}

func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.WriteError(w, appError.BadRequest("Invalied id"))
		return
	}

	post, err := h.PostService.GetByID(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, post)
}

func (h *PostHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	dto := &request.UpdatePost{}
	err := utils.ReadJSON(w, r, dto)
	if err != nil {
		utils.WriteError(w, appError.Internel(err))
		return
	}

	errs := utils.ValidateStruct(dto)
	if errs != nil {
		utils.WriteError(w, appError.BadRequest("validation error", errs))
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.WriteError(w, appError.BadRequest("Invalied id"))
		return
	}

	updatedPost, err := h.PostService.UpdateByID(r.Context(), id, dto)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, updatedPost)
}

func (h *PostHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.WriteError(w, appError.BadRequest("Invalied id"))
		return
	}

	err = h.PostService.Delete(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, &response.MessageResponce{Message: "Deleted"})
}

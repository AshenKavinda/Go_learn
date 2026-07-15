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

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: s,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	dto := &request.CreateUser{}
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

	responce, err := h.UserService.Create(ctx, dto)
	if err != nil {
		utils.WriteError(w, appError.Internel(err))
		return
	}
	utils.WriteJSON(w, http.StatusCreated, responce)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetAll(r.Context())
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.WriteError(w, appError.BadRequest("Invalid id"))
		return
	}

	user, err := h.UserService.GetByID(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	dto := &request.UpdateUser{}
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
		utils.WriteError(w, appError.BadRequest("Invalid id"))
		return
	}

	updatedUser, err := h.UserService.UpdateByID(r.Context(), id, dto)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.WriteError(w, appError.BadRequest("Invalid id"))
		return
	}

	err = h.UserService.Delete(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, &response.MessageResponce{Message: "Deleted"})
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/umeh-promise/blog/internal/controller/middlewares"
	"github.com/umeh-promise/blog/internal/services"
	"github.com/umeh-promise/blog/internal/utils"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (handler *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	user, err := handler.Service.GetUserByID(r.Context(), id)
	if err != nil {
		utils.NotFoundResponse(w, r, err)
		return
	}

	if err := utils.JSONResponse(w, http.StatusOK, user); err != nil {
		utils.InternalServerError(w, r, err)
	}
}

func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	user := middlewares.GetUserFromContext(r)

	if err := utils.JSONResponse(w, http.StatusCreated, user); err != nil {
		utils.InternalServerError(w, r, err)
	}
}

type UpdateUserPayload struct {
	FirstName    *string `json:"first_name" validate:"omitempty,max=30"`
	LastName     *string `json:"last_name" validate:"omitempty,max=30"`
	Username     *string `json:"username" validate:"omitempty,max=10"`
	ProfileImage *string `json:"profile_image"`
}

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var payload UpdateUserPayload

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	user := middlewares.GetUserFromContext(r)

	ctx := r.Context()

	if payload.FirstName != nil {
		user.FirstName = *payload.FirstName
	}
	if payload.LastName != nil {
		user.LastName = *payload.LastName
	}
	if payload.Username != nil {
		user.Username = *payload.Username
	}
	if payload.ProfileImage != nil {
		user.ProfileImage = *payload.ProfileImage
	}

	if err := handler.Service.Update(ctx, user); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	if err := utils.JSONResponse(w, http.StatusOK, user); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}
}

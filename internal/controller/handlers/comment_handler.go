package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/umeh-promise/blog/internal/controller/middlewares"
	"github.com/umeh-promise/blog/internal/models"
	"github.com/umeh-promise/blog/internal/services"
	"github.com/umeh-promise/blog/internal/utils"
)

type CommentHandler struct {
	Services *services.CommentService
}

func NewCommentHandler(services *services.CommentService) *CommentHandler {
	return &CommentHandler{Services: services}
}

type AddCommentPayload struct {
	Content string `json:"content" validate:"required"`
}

func (handler *CommentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	var payload AddCommentPayload

	user := middlewares.GetUserFromContext(r)
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	comment := &models.Comment{
		UserID:  user.ID,
		PostID:  id,
		Content: payload.Content,
		User:    *user,
	}

	if err := handler.Services.Create(r.Context(), comment); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if err := utils.JSONResponse(w, http.StatusCreated, comment); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}
}

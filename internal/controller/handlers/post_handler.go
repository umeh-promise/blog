package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/umeh-promise/blog/internal/models"
	"github.com/umeh-promise/blog/internal/services"
	"github.com/umeh-promise/blog/internal/utils"
)

type PostHandler struct {
	Service *services.PostService
}

func NewUserHandler(service *services.PostService) *PostHandler {
	return &PostHandler{
		Service: service,
	}
}

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

	var payload CreatePostPayload
	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	userID := int64(1)

	post := &models.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserID:  userID,
	}

	ctx := r.Context()

	if err := handler.Service.Create(ctx, post); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if err := utils.JSONResponse(w, http.StatusCreated, post); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}
}

func (handler *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	post, err := handler.Service.GetByID(ctx, id)
	if err != nil {
		switch err {
		case utils.ErrorNotFound:
			utils.NotFoundResponse(w, r, err)
		default:
			utils.InternalServerError(w, r, err)
		}
		return
	}

	if err := utils.JSONResponse(w, http.StatusOK, post); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

}

type UpdatePostPayload struct {
	Title   *string   `json:"title" validate:"omitempty,max=100"`
	Content *string   `json:"content" validate:"omitempty,max=1000"`
	Tags    *[]string `json:"tags"`
}

func (handler *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var payload UpdatePostPayload
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

	ctx := r.Context()

	post, err := handler.Service.Repo.GetByID(ctx, id)
	if err != nil {
		utils.NotFoundResponse(w, r, err)
		return
	}

	if payload.Title != nil {
		post.Title = *payload.Title
	}
	if payload.Content != nil {
		post.Content = *payload.Content
	}
	if payload.Tags != nil {
		post.Tags = *payload.Tags
	}

	if err := handler.Service.Repo.Update(ctx, post); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	if err := utils.JSONResponse(w, http.StatusOK, post); err != nil {
		utils.InternalServerError(w, r, err)
	}

}

func (handler *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	if err := handler.Service.Repo.Delete(r.Context(), id); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if err := utils.JSONResponse(w, http.StatusNoContent, nil); err != nil {
		utils.InternalServerError(w, r, err)
	}
}

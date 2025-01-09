package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/umeh-promise/blog/internal/controller/middlewares"
	"github.com/umeh-promise/blog/internal/models"
	"github.com/umeh-promise/blog/internal/services"
	"github.com/umeh-promise/blog/internal/utils"
)

type PostHandler struct {
	Service        *services.PostService
	CommentService *services.CommentService
}

func NewPostHandler(service *services.PostService, commentService *services.CommentService) *PostHandler {
	return &PostHandler{
		Service:        service,
		CommentService: commentService,
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

	user := middlewares.GetUserFromContext(r)

	post := &models.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserID:  user.ID,
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

	post := middlewares.GetPostFromContext(r)
	comments, err := handler.CommentService.GetCommentByPostID(r.Context(), post.ID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	post.Comments = comments

	fmt.Println(post.ID)

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

	post := middlewares.GetPostFromContext(r)
	comments, err := handler.CommentService.GetCommentByPostID(r.Context(), post.ID)
	if err != nil {
		utils.InternalServerError(w, r, err)
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

	if payload.Title != nil {
		post.Title = *payload.Title
	}
	if payload.Content != nil {
		post.Content = *payload.Content
	}
	if payload.Tags != nil {
		post.Tags = *payload.Tags
	}

	post.Comments = comments

	if err := handler.Service.Update(ctx, post); err != nil {
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

	if err := handler.Service.Delete(r.Context(), id); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if err := utils.JSONResponse(w, http.StatusNoContent, nil); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}
}

func (handler *PostHandler) GetAllPost(w http.ResponseWriter, r *http.Request) {

	postQuery := models.PostPaginationQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	postQuery, err := postQuery.Parse(r)
	if err != nil {
		utils.BadRequestError(w, r, err)
		return
	}
	posts, err := handler.Service.GetAll(r.Context(), postQuery)
	if err != nil {
		utils.NotFoundResponse(w, r, err)
		return
	}

	for i := range posts {
		comments, err := handler.CommentService.GetCommentByPostID(r.Context(), posts[i].ID)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}

		posts[i].Comments = comments
	}

	if err := utils.JSONResponse(w, http.StatusOK, posts); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}
}

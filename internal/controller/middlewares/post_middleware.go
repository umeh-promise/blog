package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/umeh-promise/blog/internal/models"
	"github.com/umeh-promise/blog/internal/services"
	"github.com/umeh-promise/blog/internal/utils"
)

type Middleware struct {
	Service *services.PostService
}

type postKey string

const postCtx postKey = "post"

func NewPostMidleware(service *services.PostService) *Middleware {
	return &Middleware{
		Service: service,
	}
}

func (middleware *Middleware) PostMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil || id <= 0 {
			utils.BadRequestError(w, r, utils.ErrorInvalidID)
			return
		}

		ctx := r.Context()

		post, err := middleware.Service.GetByID(ctx, id)
		if err != nil {
			switch err {
			case utils.ErrorNotFound:
				utils.NotFoundResponse(w, r, err)
			default:
				utils.InternalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, postCtx, post)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetPostFromContext(r *http.Request) *models.Post {
	return r.Context().Value(postCtx).(*models.Post)
}

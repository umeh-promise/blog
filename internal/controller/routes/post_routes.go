package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/umeh-promise/blog/internal/controller/handlers"
)

func PostRouter(postHandler *handlers.PostHandler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/posts", postHandler.CreatePost)
	}
}

package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/umeh-promise/blog/internal/controller/handlers"
)

func PostRouter(postHandler *handlers.PostHandler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", postHandler.CreatePost)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", postHandler.GetPostByID)
				r.Put("/", postHandler.UpdatePost)
				r.Delete("/", postHandler.DeletePost)
			})
		})
	}
}

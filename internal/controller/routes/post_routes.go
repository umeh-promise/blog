package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/umeh-promise/blog/internal/controller/handlers"
	"github.com/umeh-promise/blog/internal/controller/middlewares"
)

func PostRouter(postHandler *handlers.PostHandler, middleware *middlewares.Middleware) func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", postHandler.CreatePost)
			r.Get("/", postHandler.GetAllPost)

			r.Route("/{id}", func(r chi.Router) {
				r.Use(middleware.PostMiddleware)
				r.Get("/", postHandler.GetPostByID)
				r.Put("/", postHandler.UpdatePost)
				r.Delete("/", postHandler.DeletePost)
			})
		})
	}
}

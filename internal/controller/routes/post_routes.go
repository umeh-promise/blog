package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/umeh-promise/blog/internal/controller/handlers"
	"github.com/umeh-promise/blog/internal/controller/middlewares"
)

func PostRouter(postHandler *handlers.PostHandler, postMiddleware *middlewares.PostMiddleware, authMiddleware *middlewares.AuthMiddleware) func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", postHandler.GetAllPost)
			r.With(authMiddleware.AuthTokenMiddleware).Post("/", postHandler.CreatePost)

			r.Route("/{id}", func(r chi.Router) {
				r.Use(postMiddleware.PostMiddleware)
				r.Get("/", postHandler.GetPostByID)

				r.Group(func(r chi.Router) {
					r.Use(authMiddleware.AuthTokenMiddleware)
					r.Put("/", postHandler.UpdatePost)
					r.Delete("/", postHandler.DeletePost)
				})
			})
		})
	}
}

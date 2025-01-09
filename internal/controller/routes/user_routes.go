package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/umeh-promise/blog/internal/controller/handlers"
	"github.com/umeh-promise/blog/internal/controller/middlewares"
)

func UserRouter(userHandler *handlers.UserHandler, authMiddleware *middlewares.AuthMiddleware) func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", userHandler.Register)
			r.Post("/login", userHandler.Login)

			r.Route("/user", func(r chi.Router) {
				r.Use(authMiddleware.AuthTokenMiddleware)
				r.Get("/", userHandler.GetUser)
				r.Put("/", userHandler.UpdateUser)
			})
		})

		r.Route("/users/{id}", func(r chi.Router) {
			r.Use(authMiddleware.AuthTokenMiddleware)
			r.Get("/", userHandler.GetUserByID)
		})
	}
}

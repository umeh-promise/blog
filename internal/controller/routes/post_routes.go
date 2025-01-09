package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/umeh-promise/blog/internal/controller/handlers"
	"github.com/umeh-promise/blog/internal/controller/middlewares"
)

func PostRouter(
	postHandler *handlers.PostHandler,
	commentHandler *handlers.CommentHandler,
	postMiddleware *middlewares.PostMiddleware,
	authMiddleware *middlewares.AuthMiddleware,
	roleMiddleware *middlewares.RoleMiddleware,
) func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", postHandler.GetAllPost)
			r.With(authMiddleware.AuthTokenMiddleware).Post("/", postHandler.CreatePost)

			r.Route("/{id}", func(r chi.Router) {
				r.Use(postMiddleware.PostMiddleware)
				r.Get("/", postHandler.GetPostByID)

				r.Group(func(r chi.Router) {
					r.Use(authMiddleware.AuthTokenMiddleware)
					r.Put("/", roleMiddleware.CheckPostOwnership("moderator", postHandler.UpdatePost))
					r.Delete("/", roleMiddleware.CheckPostOwnership("admin", postHandler.DeletePost))
					r.Post("/comments", commentHandler.AddComment)
				})
			})
		})
	}
}

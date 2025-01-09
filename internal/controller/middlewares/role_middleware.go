package middlewares

import (
	"context"
	"net/http"

	"github.com/umeh-promise/blog/internal/models"
	"github.com/umeh-promise/blog/internal/services"
	"github.com/umeh-promise/blog/internal/utils"
)

type RoleMiddleware struct {
	Service *services.RoleService
}

func NewRoleMiddleware(service *services.RoleService) *RoleMiddleware {
	return &RoleMiddleware{Service: service}
}

func (middleware *RoleMiddleware) CheckPostOwnership(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := GetUserFromContext(r)
		post := GetPostFromContext(r)

		if post.UserID == user.ID {
			next.ServeHTTP(w, r)
			return
		}

		allowed, err := middleware.CheckRolePrecedence(r.Context(), user, requiredRole)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}

		if !allowed {
			utils.ForbiddenServerError(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (middleware *RoleMiddleware) CheckRolePrecedence(ctx context.Context, user *models.User, roleName string) (bool, error) {
	role, err := middleware.Service.GetByName(ctx, roleName)
	if err != nil {
		return false, err
	}

	passedUserRole, err := middleware.Service.GetByName(ctx, user.Role)
	if err != nil {
		return false, err
	}

	return passedUserRole.Level >= role.Level, nil
}

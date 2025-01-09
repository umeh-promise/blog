package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/umeh-promise/blog/internal/models"
	"github.com/umeh-promise/blog/internal/services"
	"github.com/umeh-promise/blog/internal/utils"
)

type AuthMiddleware struct {
	Service *services.UserService
}

type userKey string

var userCtx userKey = "user"

func NewAuthMiddleware(service *services.UserService) *AuthMiddleware {
	return &AuthMiddleware{Service: service}
}

func (middleware *AuthMiddleware) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.UnAuthorizedRequestError(w, r, "authorization header is missing")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.UnAuthorizedRequestError(w, r, "authorization header is malfarmed")
			return
		}

		token := parts[1]
		jwtToken, err := utils.ValidateToken(token)
		if err != nil {
			utils.UnAuthorizedRequestError(w, r, err.Error())
			return
		}

		claims := jwtToken.Claims.(jwt.MapClaims)
		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
		if err != nil {
			utils.UnAuthorizedRequestError(w, r, err.Error())
			return
		}

		ctx := r.Context()

		user, err := middleware.Service.GetUserByID(ctx, userID)
		if err != nil {
			utils.UnAuthorizedRequestError(w, r, err.Error())
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserFromContext(r *http.Request) *models.User {
	return r.Context().Value(userCtx).(*models.User)
}

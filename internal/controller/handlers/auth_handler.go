package handlers

import (
	"net/http"
	"regexp"

	"github.com/umeh-promise/blog/internal/models"
	"github.com/umeh-promise/blog/internal/utils"
)

type RegisterUserPayload struct {
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name" validate:"required,min=2"`
	LastName  string `json:"last_name" validate:"required,min=2"`
	Username  string `json:"username" validate:"required,min=2"`
	Password  string `json:"password" validate:"required,min=6"`
}

type UserWithToken struct {
	User        *models.User `json:"user"`
	AccessToken string       `json:"access_token"`
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	user := &models.User{
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Username:  payload.Username,
		Role:      "user",
	}

	if err := user.Password.Set(payload.Password); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	if err := handler.Service.Create(ctx, user); err != nil {
		switch err {
		case utils.ErrorDuplicateEmail:
			utils.BadRequestError(w, r, err)
		case utils.ErrorDuplicateUsername:
			utils.BadRequestError(w, r, err)
		default:
			utils.InternalServerError(w, r, err)
		}
		return
	}

	access_token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	userWithToken := UserWithToken{
		User:        user,
		AccessToken: access_token,
	}

	if err := utils.JSONResponse(w, http.StatusCreated, userWithToken); err != nil {
		utils.InternalServerError(w, r, err)
	}
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func isEmailVerified(email string) bool {
	return emailRegex.MatchString(email)
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginUserPayload
	if err := utils.ReadJSON(w, r, &payload); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := handler.Service.GetByEmail(ctx, payload.Email)
	if err != nil {
		switch err {
		case utils.ErrorNotFound:
			utils.UnAuthorizedRequestError(w, r, "unauthorized")
		default:
			utils.InternalServerError(w, r, err)
		}
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	userWithToken := &UserWithToken{
		User:        user,
		AccessToken: token,
	}

	if err := utils.JSONResponse(w, http.StatusOK, userWithToken); err != nil {
		utils.InternalServerError(w, r, err)
	}
}

package handlers

import (
	"net/http"

	"github.com/umeh-promise/blog/internal/services"
	"github.com/umeh-promise/blog/internal/utils"
)

type PostHandler struct {
	Service *services.PostService
}

func NewUserHandler(service *services.PostService) *PostHandler {
	return &PostHandler{
		Service: service,
	}
}

func (postHandler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"version": "1.0.0",
		"env":     "development",
	}

	if err := utils.JSONResponse(w, http.StatusOK, data); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}
}

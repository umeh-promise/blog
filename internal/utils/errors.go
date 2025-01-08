package utils

import (
	"errors"
	"net/http"

	"go.uber.org/zap"
)

var (
	ErrorNotFound  = errors.New("resource not found")
	ErrorInvalidID = errors.New("invalid post id")
	logger         = zap.Must(zap.NewProduction()).Sugar()
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	logger.Errorw("internal server error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error())

	WriteJSONError(w, http.StatusInternalServerError, []string{}, "The server encountered a problem")
}

func BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	logger.Errorw("bad request",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error())

	errors := []string{err.Error()}

	WriteJSONError(w, http.StatusBadRequest, errors, "validation errors")
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger.Errorw("not found error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error())

	WriteJSONError(w, http.StatusNotFound, []string{}, "not found")
}

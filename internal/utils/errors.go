package utils

import (
	"errors"
	"net/http"

	"go.uber.org/zap"
)

var logger = zap.Must(zap.NewProduction()).Sugar()

var (
	ErrorNotFound          = errors.New("resource not found")
	ErrorInvalidID         = errors.New("invalid post id")
	ErrorDuplicateEmail    = errors.New("a user with that email already exists")
	ErrorDuplicateUsername = errors.New("a user with that username already exists")
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	logger.Errorw("internal server error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", err.Error())

	WriteJSONError(w, http.StatusInternalServerError, []string{}, "The server encountered a problem")
}

func ForbiddenServerError(w http.ResponseWriter, r *http.Request) {
	logger.Errorw("internal server error",
		"method", r.Method,
		"path", r.URL.Path,
		"error", "forbidden")

	WriteJSONError(w, http.StatusForbidden, []string{}, "Forbidden")
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

func UnAuthorizedRequestError(w http.ResponseWriter, r *http.Request, message string) {
	logger.Errorw("unauthorized request",
		"method", r.Method,
		"path", r.URL.Path,
		"error", message)

	errors := []string{}

	WriteJSONError(w, http.StatusUnauthorized, errors, message)
}

func RateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	logger.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path)

	w.Header().Set("Retry-After", retryAfter)

	WriteJSONError(w, http.StatusTooManyRequests, []string{}, "rate limit exceeded, retry after: "+retryAfter)
}

package utils

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrorNotFound = errors.New("resource not found")
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Fatal("internal server error", err)

	WriteJSONError(w, http.StatusInternalServerError, "The server encountered a problem")
}

func BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Fatal("bad request", err)

	WriteJSONError(w, http.StatusBadRequest, err.Error())
}

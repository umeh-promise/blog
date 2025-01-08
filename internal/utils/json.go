package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1_048_576 // 1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func WriteJSONError(w http.ResponseWriter, status int, error []string, message string) error {
	type envelop struct {
		Status  string   `json:"status"`
		Error   []string `json:"error"`
		Message string   `json:"message"`
	}

	return WriteJSON(w, status, &envelop{
		Status:  "failed",
		Error:   error,
		Message: message,
	})
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) error {
	type envelop struct {
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	}

	return WriteJSON(w, status, &envelop{
		Status: "success",
		Data:   data,
	})
}

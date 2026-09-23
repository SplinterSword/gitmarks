package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func SendJson (value any, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(value); err != nil {
		SendError(errors.New("failed to encode response"), http.StatusInternalServerError, w)
		return
	}
}

func SendError (err error, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Error: err.Error(),
	}

	json.NewEncoder(w).Encode(response)
}

func ReceiveJson (schema any, w http.ResponseWriter ,r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(schema); err != nil {
		SendError(errors.New("Invalid Json"), http.StatusBadRequest, w)
		return err
	}
	return nil
}

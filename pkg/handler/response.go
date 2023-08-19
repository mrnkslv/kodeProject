package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json: "message"`
}

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	logrus.Error(message)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

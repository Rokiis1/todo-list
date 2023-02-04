package errors

import (
	"encoding/json"
	"net/http"
)

// Response is the structure of the response sent back to the client
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func WriteError(err error, w http.ResponseWriter) {
	response := Response{
		Message: "Error processing the request",
		Data:    nil,
		Error:   err.Error(),
	}

	respondWithJSON(w, http.StatusBadRequest, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

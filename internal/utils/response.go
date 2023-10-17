package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Data    interface{}
	Message string
}

func ResponseSuccess(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func RespondWithJSON(w http.ResponseWriter, data interface{}, msg string, code int) {
	jsonData := JSONResponse{
		Data:    data,
		Message: msg,
	}

	response, _ := json.Marshal(jsonData)

	w.Write(response)
}

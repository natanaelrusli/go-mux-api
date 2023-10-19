package utils

import (
	"encoding/json"
	"net/http"
)

type Pagination struct {
	Next          int `json:"next"`
	Previous      int `json:"prev"`
	RecordPerPage int `json:"perPage"`
	CurrentPage   int `json:"current"`
	TotalPage     int `json:"totalPage"`
}

type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ResponseSuccess(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func RespondWithJSON(w http.ResponseWriter, data interface{}, msg string, code int) {
	jsonData := JSONResponse{
		Success: true,
		Data:    data,
		Message: msg,
	}

	response, _ := json.Marshal(jsonData)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

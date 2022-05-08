package utils

import (
	"encoding/json"
	"net/http"
	"project/model"
)

func Return(w http.ResponseWriter, code int, err error, data interface{}) {
	status := map[int]string{
		200: "OK",
		201: "Created",
		400: "Bad Request",
		401: "Unauthorized",
		403: "Forbidden",
		404: "Not Found",
		422: "Unprocessable Entity",
		500: "Internal Server Error",
	}

	response := model.Response{
		Status: status[code],
		Code:   code,
		Error:  "",
		Data:   data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

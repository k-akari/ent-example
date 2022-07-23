package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type healthCheckResponse struct {
    Status  int    `json:"status"`
    Message string `json:"message,omitempty"`
}

func registerHealthCheckRouter(mux *http.ServeMux) {
	mux.HandleFunc("/health_check", helthCheckHandler)
}

func helthCheckHandler(w http.ResponseWriter, r *http.Request) {
	rs := healthCheckResponse{
        Status: http.StatusOK,
    }
    respondJSON(w, rs, http.StatusOK)
}

func respondJSON(w http.ResponseWriter, body interface{}, status int) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(body); err != nil {
        fmt.Fprintf(os.Stderr, "failed to encode response by error '%#v'", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}

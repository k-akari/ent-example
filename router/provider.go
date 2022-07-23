package router

import (
	"net/http"
	"project/ent"
)

func RegisterRouter(mux *http.ServeMux, c *ent.Client) {
	registerHealthCheckRouter(mux)
	registerUserRouter(mux, c)
}

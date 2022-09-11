package router

import (
	"net/http"
	"project/ent"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRouter(mux *http.ServeMux, c *ent.Client) {
	mux.Handle("/metrics", promhttp.Handler())
	registerHealthCheckRouter(mux)
	registerUserRouter(mux, c)
}

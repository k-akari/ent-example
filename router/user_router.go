package router

import (
	"context"
	"net/http"
	"path"
	"project/ent"
	"project/usecase"
)

func registerUserRouter(mux *http.ServeMux, c *ent.Client) {
	mux.HandleFunc("/users/", handleUsers(c))
}

func handleUsers(c *ent.Client) http.HandlerFunc {
	ctx := context.Background()
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if path.Base(r.URL.Path) == "users" {
				usecase.ListUsers(w, r, c, ctx)
			} else {
				usecase.ShowUser(w, r, c, ctx)
			}
		case "POST":
			usecase.CreateUser(w, r, c, ctx)
		case "PATCH":
			usecase.UpdateUser(w, r, c, ctx)
		case "DELETE":
			usecase.DeleteUser(w, r, c, ctx)
		default:
			http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

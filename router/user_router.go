package router

import (
	"context"
	"net/http"
	"path"
	"project/controller"
	"project/ent"
)

func registerUserRouter(c *ent.Client) {
	http.HandleFunc("/users/", handleUsers(c))
}

func handleUsers(c *ent.Client) http.HandlerFunc {
	ctx := context.Background()
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if path.Base(r.URL.Path) == "users" {
				controller.ListUsers(w, r, c, ctx)
			} else {
				controller.ShowUser(w, r, c, ctx)
			}
		case "POST":
			controller.CreateUser(w, r, c, ctx)
		case "PATCH":
			controller.UpdateUser(w, r, c, ctx)
		case "DELETE":
			controller.DeleteUser(w, r, c, ctx)
		default:
			http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

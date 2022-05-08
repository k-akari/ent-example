// Package controller contains handler functions
package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
	"project/ent"
	"project/ent/user"
	"project/utils"
	"strconv"
)

func HandleUsers(c *ent.Client, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if path.Base(r.URL.Path) == "users" {
				listUsers(w, r, c, ctx)
			} else {
				showUser(w, r, c, ctx)
			}
		case "POST":
			createUser(w, r, c, ctx)
		case "PATCH":
			updateUser(w, r, c, ctx)
		case "DELETE":
			deleteUser(w, r, c, ctx)
		default:
			http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func showUser(w http.ResponseWriter, r *http.Request, c *ent.Client, ctx context.Context) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		utils.Return(w, http.StatusBadRequest, err, nil)
		return
	}

	user, err := c.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		utils.Return(w, http.StatusNotFound, err, nil)
		return
	}

	utils.Return(w, http.StatusOK, nil, user)
}

func listUsers(w http.ResponseWriter, r *http.Request, c *ent.Client, ctx context.Context) {
	users, err := c.User.Query().All(ctx)
	if err != nil {
		utils.Return(w, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, http.StatusOK, nil, users)
}

func createUser(w http.ResponseWriter, r *http.Request, c *ent.Client, ctx context.Context) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var userParams *ent.User
	if err := json.Unmarshal(body, &userParams); err != nil {
		utils.Return(w, http.StatusInternalServerError, err, nil)
		return
	}

	user, err := c.User.Create().
		SetName(userParams.Name).
		SetEmail(userParams.Email).
		SetPassword(userParams.Password).
		Save(ctx)
	if err != nil {
		utils.Return(w, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, http.StatusCreated, nil, user)
}

func updateUser(w http.ResponseWriter, r *http.Request, c *ent.Client, ctx context.Context) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		utils.Return(w, http.StatusBadRequest, err, nil)
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var userParams *ent.User
	if err = json.Unmarshal(body, &userParams); err != nil {
		utils.Return(w, http.StatusInternalServerError, err, nil)
		return
	}

	user, err := c.User.UpdateOneID(id).
		SetName(userParams.Name).
		SetEmail(userParams.Email).
		Save(ctx)
	if err != nil {
		utils.Return(w, http.StatusInternalServerError, err, nil)
		return
	}

	utils.Return(w, http.StatusOK, nil, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request, c *ent.Client, ctx context.Context) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		utils.Return(w, http.StatusBadRequest, err, nil)
		return
	}

	err = c.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		utils.Return(w, http.StatusNotFound, err, nil)
		return
	}

	utils.Return(w, http.StatusOK, nil, nil)
}

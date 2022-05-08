// Package controller contains handler functions
package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"project/ent"
	"project/ent/user"
	"project/utils"
	"strconv"
)

func HandleUsers(c *ent.Client, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			if path.Base(r.URL.Path) == "users" {
				err = listUsers(w, r, c, ctx)
			} else {
				showUser(w, r, c, ctx)
			}
		case "POST":
			err = createUser(w, r, c, ctx)
		case "PATCH":
			err = updateUser(w, r, c, ctx)
		case "DELETE":
			err = deleteUser(w, r, c, ctx)
		default:
			http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(err)
			return
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

func listUsers(w http.ResponseWriter, r *http.Request, c *ent.Client, ctx context.Context) (err error) {
	users, err := c.User.Query().All(ctx)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err = enc.Encode(&users); err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, buf.String())
	return
}

func createUser(w http.ResponseWriter, r *http.Request, c *ent.Client, ctx context.Context) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var userParams *ent.User
	if err = json.Unmarshal(body, &userParams); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	user, err := c.User.Create().
		SetName(userParams.Name).
		SetEmail(userParams.Email).
		SetPassword(userParams.Password).
		Save(ctx)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err = enc.Encode(&user); err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, buf.String())
	return
}

func updateUser(w http.ResponseWriter, r *http.Request, c *ent.Client, ctx context.Context) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var userParams *ent.User
	if err = json.Unmarshal(body, &userParams); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	user, err := c.User.UpdateOneID(id).
		SetName(userParams.Name).
		SetEmail(userParams.Email).
		Save(ctx)
	if err != nil {
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))
	return
}

func deleteUser(w http.ResponseWriter, r *http.Request, c *ent.Client, ctx context.Context) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	err = c.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	return
}

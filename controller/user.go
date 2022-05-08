package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"project/ent"
	"project/ent/user"
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
				err = showUser(w, r, c, ctx)
			}
		case "POST":
			err = createUser(w, r, c, ctx)
		default:
			http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func showUser(w http.ResponseWriter, r *http.Request, c *ent.Client, ctx context.Context) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println(err)
		return
	}

	user, err := c.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))
	return
}

func listUsers(w http.ResponseWriter, r *http.Request, c *ent.Client, ctx context.Context) (err error) {
	users, err := c.User.Query().All(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&users); err != nil {
		log.Fatal(err)
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
		fmt.Println(err)
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err = enc.Encode(&user); err != nil {
		w.WriteHeader(500)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, buf.String())
	return
}

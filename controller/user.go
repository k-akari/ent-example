package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"project/ent"
)

func HandleUsers(c *ent.Client, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
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

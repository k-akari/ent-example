package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"project/ent"
	"project/router"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT") ,os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalf("failed opening connection to postgresql: %v", err)
	}
	defer client.Close()

	router.RegisterRouter(client)

	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		ReadTimeout:    time.Duration(10 * int64(time.Second)),
		WriteTimeout:   time.Duration(600 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

	// オートマイグレーションツールを実行する
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

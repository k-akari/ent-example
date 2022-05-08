package main

import (
	"context"
	"project/ent"
	"project/ent/migrate"

	"log"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=db user=user password=password dbname=database sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening postgresql client: %v", err)
	}
	defer client.Close()
	createDBSchema(client)
}

func createDBSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

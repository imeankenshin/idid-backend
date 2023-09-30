package main

import (
	"context"
	"fmt"
	"idid-api/ent"
	"idid-api/internal/settings/db"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	fmt.Println("Connected to database")
	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}
}

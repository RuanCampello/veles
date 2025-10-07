package main

import (
	"context"
	"log"

	"github.com/ruancampello/veles/internal"
)

func main() {
	ctx := context.Background()

	database, err := internal.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialise the database: %v", err)
	}

	defer database.Close()
}

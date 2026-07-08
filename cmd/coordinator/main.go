package main

import (
	"context"
	"log"
	"os"

	"github.com/imrishhh/goqueue/internal/database"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	connString := os.Getenv("DATABASE_URL")
	ctx := context.Background()
	_, err := database.NewPool(ctx, connString)
	if err != nil {
		log.Fatal("failed to establish a database pool:", err)
	}
}

package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

// We force migration to find the .env file as long as migration is ran from project directory
func init() {
	godotenv.Load()
}

func main() {
	cmd, args := "up", os.Args
	if len(args) > 2 {
		cmd = strings.ToLower(args[1])
	}
	conn, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	switch cmd {
	case "up":
		if err = goose.Up(conn, "./migrations"); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err = goose.Down(conn, "./migrations"); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Invalid command (use up to run the migration or down to remove migrations)")
	}
}

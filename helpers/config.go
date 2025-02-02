package helpers

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file", "error", err)
	}
}

func GetDBConnectionString() string {
	// Load environment variables from .env
	LoadEnv()

	// Read environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || dbname == "" {
		slog.Error("Missing required environment variables for database connection")
		slog.Info("Exiting application...")
		os.Exit(1)
	}

	// Construct the connection string
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"tutorial.sqlc.dev/app/helpers"
	"tutorial.sqlc.dev/app/tutorial"
)

func main() {
	// Initialize context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle SIGINT (Ctrl+C) for graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		<-sigChan
		cancel()
	}()

	// Create a channel to receive the connection instance
	poolChan := make(chan *pgxpool.Pool)

	// Start the run goroutine to initialize the connection instance
	go run(ctx, poolChan)

	// Wait for the connection instance from the run goroutine
	pool := <-poolChan
	defer pool.Close()

	queries := tutorial.New(pool)

	// Define flags
	createFlag := flag.Bool("c", false, "Create a new user")
	readFlag := flag.Bool("r", false, "Read a user")
	updateFlag := flag.Bool("u", false, "Update a user")
	deleteFlag := flag.Bool("d", false, "Delete a user")
	id := flag.Int64("id", 0, "The id of an author")
	name := flag.String("name", "", "The name of an author")
	bio := flag.String("bio", "", "The bio of an author")

	// Parse the flags
	flag.Parse()

	switch {
	case *createFlag:
		if *name == "" {
			slog.Error("Name is required for creating an author")
			slog.Info("Exiting application...")
			os.Exit(1)
		}
		author, err := helpers.CreateAuthor(ctx, queries, *name, *bio)
		if err != nil {
			slog.Error("Failed to create author", "error", err)
			slog.Info("Exiting application...")
			os.Exit(1)
		}
		helpers.PrintAuthor(author)
		slog.Info("Exiting application...")

	case *readFlag:
		if *id == 0 {
			authors, err := helpers.ReadAuthors(ctx, queries)
			if err != nil {
				slog.Error("Failed to read author", "error", err)
				slog.Info("Exiting application...")
				os.Exit(1)
			}
			if authors == nil {
				slog.Info("No authors present in the database")
				slog.Info("Exiting application...")
				os.Exit(1)
			} else {
				helpers.PrintAuthors(authors)
				slog.Info("Exiting application...")
			}
		} else {
			author, err := helpers.ReadAuthor(ctx, queries, *id)
			if err != nil {
				slog.Error("Failed to read author", "error", err)
				slog.Info("Exiting application...")
				os.Exit(1)
			}
			helpers.PrintAuthor(author)
			slog.Info("Exiting application...")
		}

	case *updateFlag:
		if *id == 0 {
			slog.Error("Invalid ID", "id", *id)
			slog.Info("Exiting application...")
			os.Exit(1)
		}
		if *name != "" {
			author, err := helpers.UpdateAuthorName(ctx, queries, *id, *name)
			if err != nil {
				slog.Error("Failed to update author name", "error", err)
				slog.Info("Exiting application...")
				os.Exit(1)
			}
			helpers.PrintAuthor(author)
			slog.Info("Exiting application...")
		}
		if *bio != "" {
			author, err := helpers.UpdateAuthorBio(ctx, queries, *id, *bio)
			if err != nil {
				slog.Error("Failed to update author bio", "error", err)
				slog.Info("Exiting application...")
				os.Exit(1)
			}
			helpers.PrintAuthor(author)
			slog.Info("Exiting application...")
		}

	case *deleteFlag:
		if *id == 0 {
			slog.Error("Invalid ID", "id", *id)
			slog.Info("Exiting application...")
			os.Exit(1)
		}
		if err := helpers.DeleteAuthor(ctx, queries, *id); err != nil {
			slog.Error("Failed to delete author", "error", err)
			slog.Info("Exiting application...")
			os.Exit(1)
		}
		fmt.Println("Author deleted successfully")
		slog.Info("Exiting application...")

	default:
		fmt.Println("No valid flag is set. Usage:")
		flag.PrintDefaults()
		slog.Info("Exiting application...")
	}
}

func run(ctx context.Context, c chan<- *pgxpool.Pool) {
	// Build connection string
	dbConnectionString := helpers.GetDBConnectionString()

	// Replace with your actual database connection string
	pgxPool, err := pgxpool.New(ctx, dbConnectionString)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		slog.Info("Exiting application...")
		os.Exit(1)
	}

	slog.Info("Successfully connected to the database")

	// Send the connection instance through the channel
	c <- pgxPool
}

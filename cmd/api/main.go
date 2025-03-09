package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/dudubernardino/gobank/internal/infra/api"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
	gob.Register(uuid.UUID{})
	if err := godotenv.Load(); err != nil {
		slog.Info("Warning: .env file not found, using system environment variables")
	}
}

func setupDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("GOBANK_DATABASE_USER"),
		os.Getenv("GOBANK_DATABASE_PASSWORD"),
		os.Getenv("GOBANK_DATABASE_HOST"),
		os.Getenv("GOBANK_DATABASE_PORT"),
		os.Getenv("GOBANK_DATABASE_NAME"),
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return pool, nil
}

func startServer(pool *pgxpool.Pool) error {
	api := api.Api{
		Router: chi.NewMux(),
	}

	api.BindRoutes(pool)

	serverAddr := ":3000"
	slog.Info("Starting server on", "port", serverAddr)

	return http.ListenAndServe(serverAddr, api.Router)
}

func main() {
	ctx := context.Background()

	pool, err := setupDatabase(ctx)
	if err != nil {
		slog.Error("Error initializing database", "error", err)
	}
	defer pool.Close()

	if err := startServer(pool); err != nil {
		slog.Error("Server error", "error", err)
	}
}

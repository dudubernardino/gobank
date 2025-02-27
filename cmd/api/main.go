package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"os"

	"github.com/dudubernardino/gobank/internal/infra/api"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	gob.Register(uuid.UUID{})

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", os.Getenv("GOBANK_DATABASE_USER"), os.Getenv("GOBANK_DATABASE_PASSWORD"), os.Getenv("GOBANK_DATABASE_HOST"), os.Getenv("GOBANK_DATABASE_PORT"), os.Getenv("GOBANK_DATABASE_NAME")))

	if err != nil {
		panic(err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	api := api.Api{
		Router: chi.NewMux(),
	}

	api.BindRoutes(pool)

	fmt.Println("Starting Server on port :3000")
	if err := http.ListenAndServe(":3000", api.Router); err != nil {
		panic(err)
	}
}

package api

import (
	"github.com/dudubernardino/gobank/internal/infra/api/accounts"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (api *Api) BindRoutes(pool *pgxpool.Pool) {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	api.Router.Route("/api", func(router chi.Router) {
		router.Route("/v1", func(router chi.Router) {
			router.Route("/accounts", func(router chi.Router) {
				router.Get("/{account_id}", accounts.HandleGetAccountById(pool))
			})
		})
	})
}

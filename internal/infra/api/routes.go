package api

import (
	accounts "github.com/dudubernardino/gobank/internal/infra/api/account"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (api *Api) BindRoutes(pool *pgxpool.Pool) {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	api.Router.Route("/api", func(router chi.Router) {
		router.Route("/v1", func(router chi.Router) {
			router.Route("/accounts", func(router chi.Router) {
				router.Post("/", accounts.HandleCreateAccount(pool))
				router.Get("/{account_id}", accounts.HandleGetAccountById(pool))
				router.Get("/{account_id}/balance", accounts.HandleGetAccountBalance(pool))
				router.Post("/{account_id}/deposit", accounts.HandleAccountDeposit(pool))
				router.Post("/{account_id}/withdraw", accounts.HandleAccountWithdraw(pool))
				router.Post("/transfer", accounts.HandleAccountAmountTransfer(pool))
				router.Delete("/{account_id}", accounts.HandleCloseAccount(pool))
			})
		})
	})
}

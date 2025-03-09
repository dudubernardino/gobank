package api

import (
	accounts "github.com/dudubernardino/gobank/internal/infra/api/account"
	"github.com/dudubernardino/gobank/internal/infra/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (api *Api) bindAccountRoutes(router chi.Router, pool *pgxpool.Pool) {
	accountsRepository := repositories.NewAccountsRepositoryPostgres(pool)

	router.Route("/accounts", func(router chi.Router) {
		router.Post("/", accounts.HandleCreateAccount(accountsRepository))

		router.Route("/{account_id}", func(router chi.Router) {
			router.Get("/", accounts.HandleGetAccountById(accountsRepository))
			router.Get("/balance", accounts.HandleGetAccountBalance(accountsRepository))
			router.Post("/deposit", accounts.HandleAccountDeposit(accountsRepository))
			router.Post("/withdraw", accounts.HandleAccountWithdraw(accountsRepository))
			router.Post("/transfer", accounts.HandleAccountAmountTransfer(accountsRepository))
			router.Delete("/", accounts.HandleCloseAccount(accountsRepository))
		})
	})
}

func (api *Api) BindRoutes(pool *pgxpool.Pool) {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	api.Router.Route("/api", func(router chi.Router) {
		router.Route("/v1", func(router chi.Router) {
			api.bindAccountRoutes(router, pool)
		})
	})
}

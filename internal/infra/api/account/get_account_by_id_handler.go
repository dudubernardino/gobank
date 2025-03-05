package account

import (
	"net/http"

	"github.com/dudubernardino/gobank/internal/domain/account/usecases"
	"github.com/dudubernardino/gobank/internal/infra/jsonutils"
	"github.com/dudubernardino/gobank/internal/infra/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleGetAccountById(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountRepository := repositories.NewAccountsRepository(pool)
		getAccountByIdUseCase := usecases.NewGetAccountByIdUseCase(accountRepository)

		rawAccountId := chi.URLParam(r, "account_id")
		accountId, err := uuid.Parse(rawAccountId)
		if err != nil {
			_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{"error": "invalid ID format"})
			return
		}

		account, err := getAccountByIdUseCase.Exec(usecases.GetAccountByIdUseCaseRequest{Id: accountId})

		if err != nil {
			_ = jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{"error": "account not found"})
			return
		}

		_ = jsonutils.EncodeJson(w, r, http.StatusOK, account)
	}
}

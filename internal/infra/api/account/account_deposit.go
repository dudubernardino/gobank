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

func HandleAccountDeposit(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountRepository := repositories.NewAccountsRepository(pool)
		accountDepositUseCase := usecases.NewAccountDepositUseCase(accountRepository)

		rawAccountId := chi.URLParam(r, "account_id")
		accountId, err := uuid.Parse(rawAccountId)
		if err != nil {
			_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{"error": "invalid ID format"})
			return
		}

		data, problems, err := jsonutils.DecodeJson[usecases.AccountDepositUseCaseRequest](r)
		if err != nil {
			_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, err.Error())
			return
		}

		if problems != nil {
			_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
				"validation_errors": problems,
			})
			return
		}

		balance, err := accountDepositUseCase.Exec(usecases.AccountDepositUseCaseRequest{Id: accountId, Amount: data.Amount})

		if err != nil {
			// TODO: improve this error handling
			_ = jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{"error": "account not found"})
			return
		}

		_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"balance": balance})
	}
}

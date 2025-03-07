package account

import (
	"log/slog"
	"net/http"

	"github.com/dudubernardino/gobank/internal/domain/account/usecases"
	"github.com/dudubernardino/gobank/internal/infra/jsonutils"
	"github.com/dudubernardino/gobank/internal/infra/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleAccountWithdraw(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountRepository := repositories.NewAccountsRepository(pool)
		accountWithdrawUseCase := usecases.NewAccountWithdrawUseCase(accountRepository)

		rawAccountId := chi.URLParam(r, "account_id")
		accountId, err := uuid.Parse(rawAccountId)
		if err != nil {
			_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{"error": ErrInvalidIdFormat.Error()})
			return
		}

		data, problems, err := jsonutils.DecodeJson[usecases.AccountWithdrawUseCaseRequest](r)
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

		balance, err := accountWithdrawUseCase.Exec(usecases.AccountWithdrawUseCaseRequest{Id: accountId, Amount: data.Amount})

		if err != nil {
			slog.Error("error withdrawing account balance", "error", err.Error())
			_ = jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{"error": ErrAccountWithdraw.Error()})
			return
		}

		_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"balance": balance})
	}
}

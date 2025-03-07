package account

import (
	"log/slog"
	"net/http"

	"github.com/dudubernardino/gobank/internal/domain/account/usecases"
	"github.com/dudubernardino/gobank/internal/infra/jsonutils"
	"github.com/dudubernardino/gobank/internal/infra/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleAccountAmountTransfer(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountRepository := repositories.NewAccountsRepository(pool)
		accountAmountTransferUseCase := usecases.NewAccountAmountTransferUseCase(accountRepository)

		data, problems, err := jsonutils.DecodeJson[usecases.AccountAmountTransferUseCaseRequest](r)
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

		err = accountAmountTransferUseCase.Exec(usecases.AccountAmountTransferUseCaseRequest{AccountOriginID: data.AccountOriginID, AccountDestinationID: data.AccountDestinationID, Amount: data.Amount})

		if err != nil {
			slog.Error("error transfering account balance", "data", data, "error", err.Error())
			_ = jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{"error": ErrAccountAmountTransfer.Error()})
			return
		}

		_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "amount transfered successfully"})
	}
}

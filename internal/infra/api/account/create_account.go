package account

import (
	"log/slog"
	"net/http"

	"github.com/dudubernardino/gobank/internal/domain/account/usecases"
	"github.com/dudubernardino/gobank/internal/infra/jsonutils"
	"github.com/dudubernardino/gobank/internal/infra/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleCreateAccount(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountRepository := repositories.NewAccountsRepository(pool)
		createAccountUseCase := usecases.NewCreateAccountUseCase(accountRepository)

		data, problems, err := jsonutils.DecodeJson[usecases.CreateAccountdUseCaseRequest](r)
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

		id, err := createAccountUseCase.Exec(usecases.CreateAccountdUseCaseRequest{
			TaxId:         data.TaxId,
			Name:          data.Name,
			Email:         data.Email,
			MonthlyIncome: data.MonthlyIncome,
			AnnualRevenue: data.AnnualRevenue,
			Balance:       data.Balance,
		})

		if err != nil {
			slog.Error("error creating account", "error", err.Error())
			_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{"error": ErrAccountCreate.Error()})
			return
		}

		_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{"account_id": id})
	}
}

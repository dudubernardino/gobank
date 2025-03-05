package account

import (
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

		data, err := jsonutils.DecodeJson[usecases.CreateAccountdUseCaseRequest](r)
		if err != nil {
			_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, err)
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
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{"account_id": id})
	}
}

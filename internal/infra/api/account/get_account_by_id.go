package account

import (
	"log/slog"
	"net/http"

	"github.com/dudubernardino/gobank/internal/domain/account/repositories"
	"github.com/dudubernardino/gobank/internal/domain/account/usecases"
	"github.com/dudubernardino/gobank/internal/infra/jsonutils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func HandleGetAccountById(accountRepository repositories.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getAccountByIdUseCase := usecases.NewGetAccountByIdUseCase(accountRepository)

		rawAccountId := chi.URLParam(r, "account_id")
		accountId, err := uuid.Parse(rawAccountId)
		if err != nil {
			_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{"error": ErrInvalidIdFormat.Error()})
			return
		}

		account, err := getAccountByIdUseCase.Exec(usecases.GetAccountByIdUseCaseRequest{Id: accountId})

		if err != nil {
			slog.Error("error getting account", "error", err.Error())
			_ = jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{"error": ErrAccountNotFound.Error()})
			return
		}

		_ = jsonutils.EncodeJson(w, r, http.StatusOK, account)
	}
}

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

func HandleCloseAccount(accountRepository repositories.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		closeAccountUseCase := usecases.NewCloseAccountUseCase(accountRepository)

		rawAccountId := chi.URLParam(r, "account_id")
		accountId, err := uuid.Parse(rawAccountId)
		if err != nil {
			_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{"error": ErrInvalidIdFormat.Error()})
			return
		}

		err = closeAccountUseCase.Exec(usecases.CloseAccountUseCaseRequest{Id: accountId})

		if err != nil {
			slog.Error("error closing account", "error", err.Error())
			_ = jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{"error": ErrCloseAccount.Error()})
			return
		}

		_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "account closed successfully"})
	}
}

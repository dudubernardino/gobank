package usecases

import (
	"errors"

	"github.com/dudubernardino/gobank/internal/domain/account/repositories"
	"github.com/google/uuid"
)

type AccountWithdrawUseCase struct {
	repository repositories.AccountRepository
}

func NewAccountWithdrawUseCase(repository repositories.AccountRepository) AccountWithdrawUseCase {
	return AccountWithdrawUseCase{
		repository: repository,
	}
}

type AccountWithdrawUseCaseRequest struct {
	Id     uuid.UUID `json:"id"`
	Amount int64     `json:"amount" validate:"gt=0"`
}

func (accountWithdraw *AccountWithdrawUseCase) Exec(request AccountWithdrawUseCaseRequest) (int64, error) {
	currentAccountBalance, err := accountWithdraw.repository.FindAccountBalanceById(request.Id)
	if err != nil {
		return 0, err
	}

	if currentAccountBalance < request.Amount {
		return 0, errors.New("insufficient funds")
	}

	balance, err := accountWithdraw.repository.AccountWithdraw(request.Id, request.Amount)

	if err != nil {
		return 0, err
	}

	return balance, nil
}

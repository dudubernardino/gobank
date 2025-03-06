package usecases

import (
	"github.com/dudubernardino/gobank/internal/domain/account/repositories"
	"github.com/google/uuid"
)

type AccountDepositUseCase struct {
	repository repositories.AccountRepository
}

func NewAccountDepositUseCase(repository repositories.AccountRepository) AccountDepositUseCase {
	return AccountDepositUseCase{
		repository: repository,
	}
}

type AccountDepositUseCaseRequest struct {
	Id     uuid.UUID `json:"id"`
	Amount int64     `json:"amount" validate:"gt=0"`
}

func (accountDeposit *AccountDepositUseCase) Exec(request AccountDepositUseCaseRequest) (int64, error) {
	balance, err := accountDeposit.repository.AccountDeposit(request.Id, request.Amount)

	if err != nil {
		return 0, err
	}

	return balance, nil
}

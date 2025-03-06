package usecases

import (
	"github.com/dudubernardino/gobank/internal/domain/account/repositories"
	"github.com/google/uuid"
)

type GetAccountBalanceUseCase struct {
	repository repositories.AccountRepository
}

func NewGetAccountBalanceUseCase(repository repositories.AccountRepository) GetAccountBalanceUseCase {
	return GetAccountBalanceUseCase{
		repository: repository,
	}
}

type GetAccountBalanceUseCaseRequest struct {
	Id uuid.UUID
}

func (getAccountBalance *GetAccountBalanceUseCase) Exec(request GetAccountBalanceUseCaseRequest) (int64, error) {
	balance, err := getAccountBalance.repository.FindAccountBalanceById(request.Id)

	if err != nil {
		return 0, err
	}

	return balance, nil
}

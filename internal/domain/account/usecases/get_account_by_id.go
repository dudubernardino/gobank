package usecases

import (
	"github.com/dudubernardino/gobank/internal/domain/account/entities"
	"github.com/dudubernardino/gobank/internal/domain/account/repositories"
	"github.com/google/uuid"
)

type GetAccountByIdUseCase struct {
	repository repositories.AccountRepository
}

func NewGetAccountByIdUseCase(repository repositories.AccountRepository) GetAccountByIdUseCase {
	return GetAccountByIdUseCase{
		repository: repository,
	}
}

type GetAccountByIdUseCaseRequest struct {
	Id uuid.UUID
}

func (getAccountById *GetAccountByIdUseCase) Exec(request GetAccountByIdUseCaseRequest) (entities.Account, error) {
	account, err := getAccountById.repository.FindById(request.Id)

	if err != nil {
		return entities.Account{}, err
	}

	return account, nil
}

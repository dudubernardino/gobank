package usecases

import (
	"github.com/dudubernardino/gobank/internal/domain/account/repositories"
	"github.com/google/uuid"
)

type CloseAccountUseCase struct {
	repository repositories.AccountRepository
}

func NewCloseAccountUseCase(repository repositories.AccountRepository) CloseAccountUseCase {
	return CloseAccountUseCase{
		repository: repository,
	}
}

type CloseAccountUseCaseRequest struct {
	Id uuid.UUID
}

func (closeAccount *CloseAccountUseCase) Exec(request CloseAccountUseCaseRequest) error {
	_, err := closeAccount.repository.FindById(request.Id)
	if err != nil {
		return err
	}

	err = closeAccount.repository.Delete(request.Id)
	if err != nil {
		return err
	}

	return nil
}

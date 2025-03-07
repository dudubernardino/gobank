package usecases

import (
	"errors"

	"github.com/dudubernardino/gobank/internal/domain/account/repositories"
	"github.com/google/uuid"
)

type AccountAmountTransferUseCase struct {
	repository repositories.AccountRepository
}

func NewAccountAmountTransferUseCase(repository repositories.AccountRepository) AccountAmountTransferUseCase {
	return AccountAmountTransferUseCase{
		repository: repository,
	}
}

type AccountAmountTransferUseCaseRequest struct {
	AccountOriginID      uuid.UUID `json:"account_origin_id" validate:"required"`
	AccountDestinationID uuid.UUID `json:"account_destination_id" validate:"required"`
	Amount               int64     `json:"amount" validate:"gt=0"`
}

func (accountAmountTransfer *AccountAmountTransferUseCase) Exec(request AccountAmountTransferUseCaseRequest) error {
	if request.AccountOriginID == request.AccountDestinationID {
		return errors.New("origin and destination accounts must be different")
	}

	_, err := accountAmountTransfer.repository.FindById(request.AccountOriginID)
	if err != nil {
		return err
	}

	_, err = accountAmountTransfer.repository.FindById(request.AccountDestinationID)
	if err != nil {
		return err
	}

	err = accountAmountTransfer.repository.AccountTransfer(request.AccountOriginID, request.AccountDestinationID, request.Amount)
	if err != nil {
		return err
	}

	return nil
}

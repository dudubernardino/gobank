package usecases

import (
	"github.com/dudubernardino/gobank/internal/domain/account/entities"
	"github.com/dudubernardino/gobank/internal/domain/account/repositories"
	"github.com/google/uuid"
)

type CreateAccountUseCase struct {
	repository repositories.AccountRepository
}

func NewCreateAccountUseCase(repository repositories.AccountRepository) CreateAccountUseCase {
	return CreateAccountUseCase{
		repository: repository,
	}
}

type CreateAccountdUseCaseRequest struct {
	TaxId         string `json:"tax_id" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	MonthlyIncome int64  `json:"monthly_income" validate:"gte=0"`
	AnnualRevenue int64  `json:"annual_revenue" validate:"gte=0"`
	Balance       int64  `json:"balance" validate:"gte=0"`
}

func (createAccount *CreateAccountUseCase) Exec(request CreateAccountdUseCaseRequest) (uuid.UUID, error) {
	id, err := createAccount.repository.Create(entities.Account{
		TaxID:         request.TaxId,
		Name:          request.Name,
		Email:         request.Email,
		MonthlyIncome: request.MonthlyIncome,
		AnnualRevenue: request.AnnualRevenue,
		Balance:       request.Balance,
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

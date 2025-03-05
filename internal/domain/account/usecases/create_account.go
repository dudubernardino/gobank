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
	TaxId         string `json:"tax_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	MonthlyIncome int64  `json:"monthly_income"`
	AnnualRevenue int64  `json:"annual_revenue"`
	Balance       int64  `json:"balance"`
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

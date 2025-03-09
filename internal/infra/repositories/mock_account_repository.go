package repositories

import (
	"fmt"
	"slices"
	"time"

	"github.com/dudubernardino/gobank/internal/domain/account/entities"
	"github.com/google/uuid"
)

type MockAccountRepository struct {
	Accounts []entities.Account
}

func NewMockAccountRepository(accounts []entities.Account) *MockAccountRepository {
	return &MockAccountRepository{Accounts: accounts}
}

func (mockAccountRepository *MockAccountRepository) findAccountById(id uuid.UUID) (*entities.Account, error) {
	for i, account := range mockAccountRepository.Accounts {
		if account.ID == id {
			return &mockAccountRepository.Accounts[i], nil
		}
	}
	return nil, nil
}

func (mockAccountRepository *MockAccountRepository) FindById(id uuid.UUID) (entities.Account, error) {
	account, err := mockAccountRepository.findAccountById(id)
	if err != nil {
		return *account, err
	}
	if account == nil {
		return entities.Account{}, nil
	}
	return *account, nil
}

func (mockAccountRepository *MockAccountRepository) Create(payload entities.Account) (uuid.UUID, error) {
	account := entities.Account{
		ID:            uuid.New(),
		Name:          payload.Name,
		TaxID:         payload.TaxID,
		Email:         payload.Email,
		Balance:       payload.Balance,
		MonthlyIncome: payload.MonthlyIncome,
		AnnualRevenue: payload.AnnualRevenue,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	mockAccountRepository.Accounts = append(mockAccountRepository.Accounts, account)

	return account.ID, nil
}

func (mockAccountRepository *MockAccountRepository) FindAccountBalanceById(id uuid.UUID) (int64, error) {
	account, err := mockAccountRepository.findAccountById(id)
	if err != nil {
		return 0, err
	}
	if account == nil {
		return 0, nil
	}
	return account.Balance, nil
}

func (mockAccountRepository *MockAccountRepository) AccountDeposit(id uuid.UUID, amount int64) (int64, error) {
	account, err := mockAccountRepository.findAccountById(id)
	if err != nil {
		return 0, err
	}
	if account == nil {
		return 0, nil
	}
	account.Balance += amount
	return account.Balance, nil
}

func (mockAccountRepository *MockAccountRepository) AccountWithdraw(id uuid.UUID, amount int64) (int64, error) {
	account, err := mockAccountRepository.findAccountById(id)
	if err != nil {
		return 0, err
	}
	if account == nil {
		return 0, nil
	}
	account.Balance -= amount
	return account.Balance, nil
}

func (mockAccountRepository *MockAccountRepository) AccountTransfer(originId, destinationId uuid.UUID, amount int64) error {
	originAccount, err := mockAccountRepository.findAccountById(originId)
	if err != nil || originAccount == nil {
		return fmt.Errorf("origin account not found")
	}

	if originAccount.Balance < amount {
		return fmt.Errorf("insufficient funds in origin account")
	}

	destinationAccount, err := mockAccountRepository.findAccountById(destinationId)
	if err != nil || destinationAccount == nil {
		return fmt.Errorf("destination account not found")
	}

	originAccount.Balance -= amount
	destinationAccount.Balance += amount

	return nil
}

func (mockAccountRepository *MockAccountRepository) Delete(id uuid.UUID) error {
	account, err := mockAccountRepository.findAccountById(id)
	if err != nil {
		return fmt.Errorf("error finding account: %v", err)
	}
	if account == nil {
		return fmt.Errorf("account not found")
	}

	for i, account := range mockAccountRepository.Accounts {
		if account.ID == id {
			mockAccountRepository.Accounts = slices.Delete(mockAccountRepository.Accounts, i, i+1)
			return nil
		}
	}

	return fmt.Errorf("account not found")
}

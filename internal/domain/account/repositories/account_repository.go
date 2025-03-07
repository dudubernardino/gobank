package repositories

import (
	"github.com/dudubernardino/gobank/internal/domain/account/entities"
	"github.com/google/uuid"
)

type AccountRepository interface {
	FindById(id uuid.UUID) (entities.Account, error)
	Create(account entities.Account) (uuid.UUID, error)
	FindAccountBalanceById(id uuid.UUID) (int64, error)
	AccountDeposit(id uuid.UUID, amount int64) (int64, error)
	AccountWithdraw(id uuid.UUID, amount int64) (int64, error)
	AccountTransfer(originId, destinationId uuid.UUID, amount int64) error
	Delete(id uuid.UUID) error
}
